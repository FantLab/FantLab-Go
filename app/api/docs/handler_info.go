package docs

import (
	"fantlab/api/internal/routing"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"

	"golang.org/x/tools/go/packages"

	_ "fantlab/pb"
)

var td types.TypeAndValue

func getEndpointsInfo(endpoints []routing.Endpoint) []string {
	var pkg *packages.Package

	for _, endpoint := range endpoints {
		frame := getCallerFrame(endpoint.Handler())
		if frame == nil {
			continue
		}

		dir := filepath.Dir(frame.File)
		if pkg == nil {
			pkg = loadPackage(dir)
		}

		for _, f := range pkg.Syntax {
			ast.Inspect(f, func(node ast.Node) bool {
				if fn, ok := node.(*ast.FuncDecl); ok {
					if strings.Contains(frame.Function, fn.Name.Name) {
						fmt.Println(fn.Name.Name)
						fmt.Println(getFuncComment(f, pkg.Fset, frame.Line))
						printPossibleReturnTypes(fn, pkg)
						fmt.Println()

						return false
					}
				}

				return true
			})
		}

	}

	return nil
}

func printPossibleReturnTypes(fnDecl *ast.FuncDecl, pkg *packages.Package) {
	ast.Inspect(fnDecl, func(node ast.Node) bool {
		if returnStmt, ok := node.(*ast.ReturnStmt); ok {
			results := returnStmt.Results

			if len(results) != 2 {
				return false
			}

			code := getStatusCode(results[0], pkg)
			message := getProtoMessageInstance(results[1], pkg)

			fmt.Println(code, getProtoMessageJSONWithDefaults(message))

			return false
		}

		return true
	})
}

func getProtoMessageJSONWithDefaults(pb proto.Message) string {
	return proto.MarshalTextString(pb)
	// m := jsonpb.Marshaler{
	// 	OrigName:     true,
	// 	EmitDefaults: true,
	// 	Indent:       "  ",
	// }

	// sb := new(strings.Builder)

	// _ = m.Marshal(sb, pb)

	// return sb.String()
}

func getStatusCode(expr ast.Expr, pkg *packages.Package) (code int) {
	t := pkg.TypesInfo.Types[expr]

	if t != td {
		code, _ = strconv.Atoi(t.Value.ExactString())
	}

	return
}

func getProtoMessageInstance(expr ast.Expr, pkg *packages.Package) (message proto.Message) {
	t := pkg.TypesInfo.Types[expr]

	if t != td {
		name := getProtoNameFromTypeName(t.Type.String())
		messageType := proto.MessageType(name).Elem()
		message = reflect.New(messageType).Interface().(proto.Message)
	}

	return
}

func getProtoNameFromTypeName(typeName string) string {
	typeName = typeName[(strings.Index(typeName, ".") + 1):]

	return strings.ReplaceAll(typeName, "_", ".")
}

func loadPackage(dir string) *packages.Package {
	cfg := &packages.Config{
		Mode: packages.LoadSyntax,
	}

	pkgs, err := packages.Load(cfg, dir)

	if err != nil {
		return nil
	}

	return pkgs[0]
}

func getCallerFrame(i interface{}) *runtime.Frame {
	pc := reflect.ValueOf(i).Pointer()
	frames := runtime.CallersFrames([]uintptr{pc})
	if frames == nil {
		return nil
	}
	frame, _ := frames.Next()
	if frame.Entry == 0 {
		return nil
	}
	return &frame
}

func getFuncComment(file *ast.File, fset *token.FileSet, line int) string {
	for _, cmt := range file.Comments {
		if fset.Position(cmt.End()).Line+1 == line {
			return cmt.Text()
		}
	}

	return ""
}
