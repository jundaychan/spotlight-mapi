package report

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// TestNoUnguardedEmbeddedUnmarshaler 静态守卫：扫全仓库，凡是**匿名内嵌了一个自带
// UnmarshalJSON 的类型**、自己却没有 UnmarshalJSON 的结构体，一律判失败。
//
// 这类结构体在标准 json.Unmarshal 下会静默丢掉自己声明的全部字段（见 embed.go 的
// 长注释）。修法是加四行：
//
//	func (r *T) UnmarshalJSON(b []byte) error { return report.UnmarshalEmbedded(b, r) }
//
// 有这条守卫，以后新加报表类型忘了这一步会当场红，而不是等到线上某个字段恒为空。
func TestNoUnguardedEmbeddedUnmarshaler(t *testing.T) {
	root := repoRoot(t)
	fset := token.NewFileSet()

	// pkgName.TypeName -> 是否自带 UnmarshalJSON
	hasUnmarshal := map[string]bool{}
	// 待检查：pkgName.TypeName -> 它内嵌的那些 pkg.Type
	embeds := map[string][]string{}
	where := map[string]string{}

	walkGoFiles(t, root, func(path string) {
		f, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			t.Fatalf("parse %s: %v", path, err)
		}
		pkg := f.Name.Name

		for _, d := range f.Decls {
			// 方法：收集谁有 UnmarshalJSON
			if fn, ok := d.(*ast.FuncDecl); ok && fn.Recv != nil && fn.Name.Name == "UnmarshalJSON" {
				if n := recvTypeName(fn.Recv); n != "" {
					hasUnmarshal[pkg+"."+n] = true
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
				for _, fld := range st.Fields.List {
					if len(fld.Names) != 0 { // 具名字段不受影响，方法不会被提升
						continue
					}
					if n := embeddedTypeName(fld.Type, pkg); n != "" {
						embeds[key] = append(embeds[key], n)
						where[key] = path
					}
				}
			}
		}
	})

	var bad []string
	for outer, list := range embeds {
		if hasUnmarshal[outer] {
			continue // 自己实现了，安全
		}
		for _, inner := range list {
			if hasUnmarshal[inner] {
				rel, _ := filepath.Rel(root, where[outer])
				bad = append(bad, outer+" 内嵌了 "+inner+"（"+rel+"）")
			}
		}
	}
	sort.Strings(bad)
	if len(bad) > 0 {
		t.Errorf("以下类型匿名内嵌了自带 UnmarshalJSON 的类型，却没有自己的 UnmarshalJSON，\n"+
			"标准 json 解码会静默丢掉它们自己声明的全部字段。\n"+
			"修法：func (r *T) UnmarshalJSON(b []byte) error { return report.UnmarshalEmbedded(b, r) }\n  %s",
			strings.Join(bad, "\n  "))
	}
}

func repoRoot(t *testing.T) string {
	t.Helper()
	// 本文件在 <root>/model/report/
	root, err := filepath.Abs("../..")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, "go.mod")); err != nil {
		t.Fatalf("没找到仓库根（%s）: %v", root, err)
	}
	return root
}

func walkGoFiles(t *testing.T, root string, fn func(path string)) {
	t.Helper()
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if name := info.Name(); name != "." && strings.HasPrefix(name, ".") || name == "testdata" {
				if path != root {
					return filepath.SkipDir
				}
			}
			return nil
		}
		if strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go") {
			fn(path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func recvTypeName(recv *ast.FieldList) string {
	if len(recv.List) == 0 {
		return ""
	}
	switch e := recv.List[0].Type.(type) {
	case *ast.StarExpr:
		if id, ok := e.X.(*ast.Ident); ok {
			return id.Name
		}
	case *ast.Ident:
		return e.Name
	}
	return ""
}

// embeddedTypeName 把匿名内嵌字段的类型解析成 "pkg.Type"（同包则补上当前包名）。
func embeddedTypeName(e ast.Expr, curPkg string) string {
	switch v := e.(type) {
	case *ast.StarExpr:
		return embeddedTypeName(v.X, curPkg)
	case *ast.Ident:
		return curPkg + "." + v.Name
	case *ast.SelectorExpr:
		if id, ok := v.X.(*ast.Ident); ok {
			return id.Name + "." + v.Sel.Name
		}
	}
	return ""
}
