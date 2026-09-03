package report

import (
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"
	"sort"
	"strings"
	"testing"
)

// 已人工核对：注释里的 "0-" 是 "10-xxx"/"20-xxx" 这类枚举值的一部分，或 0 本身不是合法取值。
// 加白之前请先确认：**这个字段的 0 有语义吗？聚光不传时的默认是不是 0？**
// 两个答案只要不一致，就该改成指针，而不是加到这里。
var omitemptyReviewedOK = map[string]bool{
	"campaign.ListRequest.marketing_target":                 true, // 3/4/8/9/10/13/14/15/16，无 0
	"newcreate.CreateCampaign.marketing_target":             true, // 4/9/13/16/20/21，无 0
	"offline.Request.marketing_target":                      true, // 同上
	"simpledelivery.UbeSemiBaseConfigDTO.marketing_target":  true, // 16/20/21，无 0
	"simpledelivery.UbeSemiBaseConfigDTO.carrier_type":      true, // 4~10/13，无 0
	"simpledelivery.UbeSemiBaseConfigDTO.constraint_value":  true, // 出价 10~499999 分，0 非法
	"simpledelivery.TargetRealtimeRequest.marketing_target": true,
	"creativity.StatusUpdateRequest.action_type":            true, // 1/2/3，无 0
	"finance.UpdateDailyBudgetRequest.account_budget":       true, // 金额，0 非法
	"unit.TargetConfig.keyword_target_period":               true, // 3/7/15/30，无 0
}

var zeroEnumSignal = regexp.MustCompile(`0[-=－]|false-|默认\s*-1|默认值\s*-1|默认值-1`)

// TestNoPlainZeroEnumInRequests 静态守卫：会被序列化发出去的请求结构体里，
// 「0 是合法枚举值」的字段不许用裸 int/bool + omitempty——零值会被 json 静默丢掉。
//
// 踩过两次：一次是新创编（optimize_objective=0 点击量、phrase_match_type=0 精确匹配
// 等九个字段整批发不出去，代码里写着却没生效）；一次是有人干脆在业务侧手写本地结构体
// 绕开 SDK（注释就是「必填，0 也要发」）。修法是把字段改成 *int/*bool，
// 赋值用 util.Ptr(0)。
func TestNoPlainZeroEnumInRequests(t *testing.T) {
	root := repoRoot(t)
	fset := token.NewFileSet()

	type field struct {
		tag, doc, typ string
	}
	structs := map[string][]field{} // pkg.Type -> 字段
	refs := map[string][]string{}   // pkg.Type -> 引用到的 pkg.Type
	encoders := map[string]bool{}   // 有 Encode() 的请求根类型

	walkGoFiles(t, root, func(path string) {
		f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			t.Fatalf("parse %s: %v", path, err)
		}
		pkg := f.Name.Name
		for _, d := range f.Decls {
			if fn, ok := d.(*ast.FuncDecl); ok && fn.Recv != nil && fn.Name.Name == "Encode" {
				if n := recvTypeName(fn.Recv); n != "" {
					encoders[pkg+"."+n] = true
				}
				continue
			}
			gd, ok := d.(*ast.GenDecl)
			if !ok || gd.Tok != token.TYPE {
				continue
			}
			for _, spec := range gd.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				st, ok := ts.Type.(*ast.StructType)
				if !ok || st.Fields == nil {
					continue
				}
				key := pkg + "." + ts.Name.Name
				for _, fl := range st.Fields.List {
					typ, isPtr := typeString(fl.Type, pkg)
					if typ != "" {
						refs[key] = append(refs[key], typ)
					}
					if fl.Tag == nil {
						continue
					}
					tag := strings.Trim(fl.Tag.Value, "`")
					name, _, _ := strings.Cut(strings.TrimPrefix(extractJSON(tag), ""), ",")
					if name == "" || name == "-" || !strings.Contains(tag, "omitempty") {
						continue
					}
					base, _ := typeString(fl.Type, pkg)
					if isPtr || !isPlainScalar(fl.Type) {
						continue
					}
					_ = base
					doc := ""
					if fl.Doc != nil {
						doc = fl.Doc.Text()
					}
					structs[key] = append(structs[key], field{tag: name, doc: doc, typ: "scalar"})
				}
			}
		}
	})

	// 从有 Encode() 的类型出发做可达性：只有真会被发出去的结构体才算数
	reach := map[string]bool{}
	var q []string
	for e := range encoders {
		q = append(q, e)
	}
	for len(q) > 0 {
		cur := q[len(q)-1]
		q = q[:len(q)-1]
		if reach[cur] {
			continue
		}
		reach[cur] = true
		for _, r := range refs[cur] {
			if !reach[r] {
				q = append(q, r)
			}
		}
	}

	var bad []string
	for typ, fs := range structs {
		if !reach[typ] {
			continue // 响应结构体：omitempty 只影响序列化，解码不受影响
		}
		for _, f := range fs {
			if !zeroEnumSignal.MatchString(f.doc) {
				continue
			}
			if omitemptyReviewedOK[typ+"."+f.tag] {
				continue
			}
			bad = append(bad, typ+"."+f.tag+"  // "+strings.TrimSpace(strings.ReplaceAll(f.doc, "\n", " ")))
		}
	}
	sort.Strings(bad)
	if len(bad) > 0 {
		t.Errorf("以下请求字段的 0/false 是合法枚举值，却用了裸 int/bool + omitempty，\n"+
			"零值会被 json 静默丢掉、发不到聚光。改成 *int/*bool（赋值用 util.Ptr(0)），\n"+
			"确认 0 不是合法取值的加进 omitemptyReviewedOK 并写明理由。\n  %s",
			strings.Join(bad, "\n  "))
	}
}

func isPlainScalar(e ast.Expr) bool {
	id, ok := e.(*ast.Ident)
	if !ok {
		return false
	}
	switch id.Name {
	case "int", "int64", "int32", "uint64", "bool":
		return true
	}
	return false
}

// typeString 把字段类型解析成 "pkg.Type"（切片/指针剥壳），非具名类型返回 ""。
func typeString(e ast.Expr, curPkg string) (string, bool) {
	switch v := e.(type) {
	case *ast.StarExpr:
		s, _ := typeString(v.X, curPkg)
		return s, true
	case *ast.ArrayType:
		return typeString(v.Elt, curPkg)
	case *ast.Ident:
		if isPlainScalar(v) || v.Name == "string" || v.Name == "float64" || v.Name == "any" {
			return "", false
		}
		return curPkg + "." + v.Name, false
	case *ast.SelectorExpr:
		if id, ok := v.X.(*ast.Ident); ok {
			return id.Name + "." + v.Sel.Name, false
		}
	}
	return "", false
}

func extractJSON(tag string) string {
	i := strings.Index(tag, `json:"`)
	if i < 0 {
		return ""
	}
	rest := tag[i+6:]
	j := strings.Index(rest, `"`)
	if j < 0 {
		return ""
	}
	return rest[:j]
}
