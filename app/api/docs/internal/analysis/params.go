package analysis

import (
	"go/ast"
	"go/types"
	"reflect"
	"strings"

	"golang.org/x/tools/go/packages"
)

func getRequestParams(funcDecl *ast.FuncDecl, pkg *packages.Package) []*ParamInfo {
	var paramsType types.Type

	ast.Inspect(funcDecl, func(node ast.Node) bool {
		if call, ok := node.(*ast.CallExpr); ok {
			if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
				if sel.Sel.Name == "bindParams" {
					if u, ok := call.Args[0].(*ast.UnaryExpr); ok {
						paramsType = pkg.TypesInfo.Types[u.X].Type
					}
				}
			}
		}
		return paramsType == nil
	})

	if paramsType == nil {
		return nil
	}

	var paramsStruct *ast.StructType

	ast.Inspect(funcDecl, func(node ast.Node) bool {
		if strct, ok := node.(*ast.StructType); ok {
			if pkg.TypesInfo.Types[strct].Type == paramsType {
				paramsStruct = strct
			}
		}
		return paramsStruct == nil
	})

	if paramsStruct == nil {
		return nil
	}

	var params []*ParamInfo

	ast.Inspect(paramsStruct, func(node ast.Node) bool {
		if flist, ok := node.(*ast.FieldList); ok {
			for i := 0; i < flist.NumFields(); i++ {
				f := flist.List[i]
				p := getParamInfo(f, pkg)
				if p != nil {
					params = append(params, p)
				}
			}
			return false
		}
		return true
	})

	return params
}

func getParamInfo(f *ast.Field, pkg *packages.Package) *ParamInfo {
	typeName := nameOfType(pkg.TypesInfo.Types[f.Type].Type)
	if "" == typeName {
		return nil
	}

	name, source := getNameAndSource(f.Tag.Value)
	if "" == name || !isValidSource(source) {
		return nil
	}

	return &ParamInfo{
		Name:        name,
		Source:      source,
		TypeName:    typeName,
		Description: f.Doc.Text(),
	}
}

var validParamKinds = []types.BasicKind{
	types.Bool,
	types.Int,
	types.Int8,
	types.Int16,
	types.Int32,
	types.Int64,
	types.Uint,
	types.Uint8,
	types.Uint16,
	types.Uint32,
	types.Uint64,
	types.String,
}

func nameOfType(t types.Type) string {
	if t != nil {
		if bt, ok := t.(*types.Basic); ok {
			for _, k := range validParamKinds {
				if bt.Kind() == k {
					return bt.Name()
				}
			}
		}
	}
	return ""
}

func getNameAndSource(tag string) (name string, source string) {
	n := len(tag)
	if n < 12 { // тег не должен быть короче, чем `http:"x:y"`
		return
	}

	tag = reflect.StructTag(tag[1 : n-1]).Get("http")

	tags := strings.Split(tag, ",")
	if len(tags) != 2 {
		return
	}

	return strings.TrimSpace(tags[0]), strings.TrimSpace(tags[1])
}

func isValidSource(source string) bool {
	return source == "path" || source == "query" || source == "form"
}
