package report

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"sort"
	"strings"
	"testing"
)

// TestNoScalarOmitemptyInResponses 静态守卫：只用于解码的响应结构体，数值/布尔字段不许带 omitempty。
//
// 解码时 omitempty 不起作用，看起来无害；但业务侧常把 SDK 响应原样再序列化给前端，
// 这时有意义的 0 会被整字段吃掉——前端拿到的是「没这个字段」而不是 0。
// 踩过：账户余额为 0 的 6 个账户在余额页显示成「上游没返回余额金额」，
// 恰恰是最该提醒充值的那几个被藏起来了。
//
// 请求结构体（从某个 Encode() 可达）不归这里管，那边的规矩见 TestNoPlainZeroEnumInRequests。
// 本地修正：FIX=1 go test -run TestNoScalarOmitemptyInResponses ./model/report/
func TestNoScalarOmitemptyInResponses(t *testing.T) {
	root := repoRoot(t)
	fset := token.NewFileSet()

	type hit struct {
		typ, tag string
		pos      token.Position
	}
	fields := map[string][]hit{}  // pkg.Type -> 可疑字段
	refs := map[string][]string{} // pkg.Type -> 引用的 pkg.Type
	encoders := map[string]bool{}
	scalarNamed := map[string]bool{}
	scalar := func(n string) bool {
		switch n {
		case "int", "int64", "int32", "uint64", "float64", "bool":
			return true
		}
		return false
	}

	walkGoFiles(t, root, func(path string) {
		f, err := parser.ParseFile(fset, path, nil, 0)
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
				if id, ok := ts.Type.(*ast.Ident); ok && scalar(id.Name) {
					scalarNamed[pkg+"."+ts.Name.Name] = true
					continue
				}
				st, ok := ts.Type.(*ast.StructType)
				if !ok || st.Fields == nil {
					continue
				}
				key := pkg + "." + ts.Name.Name
				for _, fl := range st.Fields.List {
					named, isPtr := typeString(fl.Type, pkg)
					if named != "" {
						refs[key] = append(refs[key], named)
					}
					if fl.Tag == nil || isPtr {
						continue
					}
					if _, isSlice := fl.Type.(*ast.ArrayType); isSlice {
						continue
					}
					tag := strings.Trim(fl.Tag.Value, "`")
					js := extractJSON(tag)
					name, opts, _ := strings.Cut(js, ",")
					if name == "" || name == "-" || !strings.Contains(opts, "omitempty") {
						continue
					}
					id, isIdent := fl.Type.(*ast.Ident)
					plain := isIdent && scalar(id.Name)
					if !plain && named == "" {
						continue
					}
					fields[key] = append(fields[key], hit{typ: named, tag: name, pos: fset.Position(fl.Tag.Pos())})
				}
			}
		}
	})

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
		q = append(q, refs[cur]...)
	}

	var bad []hit
	for typ, hs := range fields {
		if reach[typ] {
			continue
		}
		for _, h := range hs {
			if h.typ != "" && !scalarNamed[h.typ] {
				continue
			}
			h.typ = typ
			bad = append(bad, h)
		}
	}
	sort.Slice(bad, func(i, j int) bool {
		if bad[i].pos.Filename != bad[j].pos.Filename {
			return bad[i].pos.Filename < bad[j].pos.Filename
		}
		return bad[i].pos.Line < bad[j].pos.Line
	})
	if len(bad) == 0 {
		return
	}

	if os.Getenv("FIX") == "1" {
		byFile := map[string][]int{}
		for _, h := range bad {
			byFile[h.pos.Filename] = append(byFile[h.pos.Filename], h.pos.Line)
		}
		for file, lines := range byFile {
			raw, err := os.ReadFile(file)
			if err != nil {
				t.Fatal(err)
			}
			ls := strings.Split(string(raw), "\n")
			for _, n := range lines {
				ls[n-1] = strings.Replace(ls[n-1], `,omitempty"`, `"`, 1)
			}
			if err := os.WriteFile(file, []byte(strings.Join(ls, "\n")), 0o644); err != nil {
				t.Fatal(err)
			}
		}
		t.Logf("已修正 %d 个字段（%d 个文件）", len(bad), len(byFile))
		return
	}

	out := make([]string, 0, len(bad))
	for _, h := range bad {
		out = append(out, fmt.Sprintf("%s:%d  %s.%s", h.pos.Filename, h.pos.Line, h.typ, h.tag))
	}
	t.Errorf("以下响应字段是数值/布尔 + omitempty：业务侧再序列化时有意义的 0/false 会被整字段吃掉。\n"+
		"去掉 omitempty（FIX=1 可自动修正）。\n  %s", strings.Join(out, "\n  "))
}
