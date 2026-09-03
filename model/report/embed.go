package report

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync"
)

// UnmarshalEmbedded 解决 Go 的一个静默陷阱：结构体**匿名内嵌**了一个自带
// UnmarshalJSON 的类型（本包的 DataReportDTO 就是，它用来兼容 ube/* 的 camelCase 指标键）
// 时，那个方法会被提升到外层，于是 json.Unmarshal 走的是内嵌类型的解码逻辑，
// **外层结构体自己声明的字段一个都不会被填**——不报错、字段数对得上，就是值全是零值。
//
// 曾经的代价：reportoffline2.NoteReport 的 note_id 恒为空串，笔记线索同步表面
// "成功"实际一行没入库，被误判成"这个账号历史真的没数据"，2026-09-03 在生产环境
// 造成过一次数据丢失。
//
// 用法：给外层类型加四行，不用手抄字段——
//
//	func (r *SearchWordReport) UnmarshalJSON(b []byte) error {
//		return report.UnmarshalEmbedded(b, r)
//	}
//
// 做法：反射出一份"把内嵌层摊平、且不含任何匿名字段"的镜像类型（所以不会再继承
// 任何 UnmarshalJSON），用它解码外层自有字段；匿名内嵌里**自带 UnmarshalJSON 的**
// 那些则原样委托给它们自己的方法。两边都不丢。
//
// 摊平时按 encoding/json 的规则处理同名冲突：层级浅的赢；json:"-" 与未导出字段跳过。
// 解码前会把 outer 的当前值先灌进镜像，所以报文里没出现的字段保持原值，
// 与标准库行为一致。
func UnmarshalEmbedded(data []byte, outer any) error {
	rv := reflect.ValueOf(outer)
	if rv.Kind() != reflect.Ptr || rv.IsNil() || rv.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("report.UnmarshalEmbedded: outer 必须是非 nil 的结构体指针，收到 %T", outer)
	}
	ev := rv.Elem()
	l := layoutFor(ev.Type())

	// 1) 自带 UnmarshalJSON 的内嵌层，交回给它自己解（DataReportDTO 的 camelCase 兼容在这里生效）
	for _, path := range l.delegates {
		u, ok := ev.FieldByIndex(path).Addr().Interface().(json.Unmarshaler)
		if !ok {
			continue
		}
		if err := u.UnmarshalJSON(data); err != nil {
			return err
		}
	}

	// 2) 其余字段用镜像类型解——镜像里没有匿名字段，不会再继承任何 UnmarshalJSON
	mv := reflect.New(l.mirror).Elem()
	for i, path := range l.paths {
		mv.Field(i).Set(ev.FieldByIndex(path))
	}
	if err := json.Unmarshal(data, mv.Addr().Interface()); err != nil {
		return err
	}
	for i, path := range l.paths {
		ev.FieldByIndex(path).Set(mv.Field(i))
	}
	return nil
}

type embedLayout struct {
	mirror    reflect.Type
	paths     [][]int // 镜像第 i 个字段 → outer 里的下标路径
	delegates [][]int // 自带 UnmarshalJSON 的匿名内嵌字段路径
}

var (
	layoutCache     sync.Map // reflect.Type -> *embedLayout
	unmarshalerType = reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
)

func layoutFor(t reflect.Type) *embedLayout {
	if v, ok := layoutCache.Load(t); ok {
		return v.(*embedLayout)
	}
	l := buildLayout(t)
	layoutCache.Store(t, l)
	return l
}

func buildLayout(t reflect.Type) *embedLayout {
	l := &embedLayout{}
	taken := make(map[string]bool)
	var mirrorFields []reflect.StructField

	type node struct {
		t    reflect.Type
		path []int
	}
	// 按层广度遍历：encoding/json 的同名冲突规则是"层级浅的赢"
	level := []node{{t, nil}}
	for len(level) > 0 {
		var next []node
		type cand struct {
			f    reflect.StructField
			path []int
			name string
		}
		var cands []cand

		for _, nd := range level {
			for i := 0; i < nd.t.NumField(); i++ {
				f := nd.t.Field(i)
				path := append(append([]int(nil), nd.path...), i)

				if f.Anonymous && f.Type.Kind() == reflect.Struct {
					if reflect.PointerTo(f.Type).Implements(unmarshalerType) {
						l.delegates = append(l.delegates, path)
						continue // 它自己解自己，不进镜像
					}
					next = append(next, node{f.Type, path}) // 普通内嵌：下一层摊平
					continue
				}
				if f.PkgPath != "" { // 未导出，json 本来也碰不到
					continue
				}
				name, _, _ := strings.Cut(f.Tag.Get("json"), ",")
				if name == "-" {
					continue
				}
				if name == "" {
					name = f.Name
				}
				cands = append(cands, cand{f, path, name})
			}
		}
		for _, c := range cands {
			if taken[c.name] {
				continue // 浅层已经占了这个 json 名
			}
			taken[c.name] = true
			mirrorFields = append(mirrorFields, reflect.StructField{
				Name: fmt.Sprintf("F%d", len(mirrorFields)), // 必须导出，reflect.StructOf 才允许赋值
				Type: c.f.Type,
				Tag:  c.f.Tag,
			})
			l.paths = append(l.paths, c.path)
		}
		level = next
	}
	l.mirror = reflect.StructOf(mirrorFields)
	return l
}
