package generator

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"

	"github.com/BiaoLiu/go-i18n/cmd/i18n/internal/utils"
)

func Extract(extractPath string, outputPath string) error {
	data := make(map[string]string)
	files, err := utils.GetFilesBySuffix(extractPath, ".go")
	if err != nil {
		return err
	}
	for _, name := range files {
		fmt.Println(name)
		content, err := os.ReadFile(name)
		if err != nil {
			return err
		}
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, name, string(content), 0)
		if err != nil {
			return err
		}

		ast.Inspect(f, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			fn, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			pack, ok := fn.X.(*ast.Ident)
			if !ok {
				return true
			}
			if pack.Name != "i18n" {
				return true
			}
			if len(call.Args) == 0 {
				return true
			}

			var expr ast.Expr
			// if Fprintf, we'll take second arg as template
			if fn.Sel.Name == "Fprintf" {
				expr = call.Args[1]
			} else { // include Printf, Sprintf
				expr = call.Args[0]
			}
			str, ok := expr.(*ast.BasicLit)
			if !ok {
				return true
			}
			// Keep this for later debug usage.
			// log.Printf("%v", str.Value)
			data[str.Value] = str.Value
			return true
		})
	}
	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	// fmt.Printf("🚀 Creating %s i18n code project... \n\n", outputPath)
	if err := utils.CreateFolder(filepath.Join(outputPath, "translations", "en_US")); err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(outputPath, "/translations/en_US/data.json"), content, 0664)
	if err != nil {
		return err
	}
	return nil
}
