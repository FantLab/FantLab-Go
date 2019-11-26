package docs

import (
	"fantlab/api/internal/routing"
	"fantlab/pb"
	"fmt"
	"go/ast"
	"go/token"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"

	"golang.org/x/tools/go/packages"

	_ "fantlab/pb"
)

func getEndpointsInfo(endpoints []routing.Endpoint) (result []string) {
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
				if funcDecl, ok := node.(*ast.FuncDecl); ok {
					if strings.Contains(frame.Function, funcDecl.Name.Name) {
						comment := getFuncComment(f, pkg.Fset, frame.Line)
						scheme := getEndpointResponseScheme(funcDecl, pkg)
						result = append(result, comment+"\n\n"+scheme+"\n\n")
						return false
					}
				}
				return true
			})
		}
	}
	return
}

func getEndpointResponseScheme(funcDecl *ast.FuncDecl, pkg *packages.Package) (result string) {
	ast.Inspect(funcDecl, func(node ast.Node) bool {
		if returnStmt, ok := node.(*ast.ReturnStmt); ok {
			var statusCode int
			var messageType reflect.Type

			for i, expr := range returnStmt.Results {
				if t, ok := pkg.TypesInfo.Types[expr]; ok {
					switch i {
					case 0:
						statusCode, _ = strconv.Atoi(t.Value.ExactString())
					case 1:
						protoName := getProtoNameFromTypeName(t.Type.String())
						messageType = proto.MessageType(protoName)
					default:
						break
					}
				}
			}

			if statusCode > 0 && messageType != nil && messageType != reflect.TypeOf((*pb.Error_Response)(nil)) {
				fr := getCallerFrame(reflect.New(messageType.Elem()).Interface())

				if fr != nil {
					fmt.Println(fr.File)
				}
				result = getScheme(messageType, "  ")
			}

			return false
		}

		return "" == result
	})
	return
}

func getProtoNameFromTypeName(typeName string) string {
	typeName = typeName[(strings.Index(typeName, ".") + 1):]
	return strings.ReplaceAll(typeName, "_", ".")
}

func loadPackage(dir string) *packages.Package {
	cfg := &packages.Config{
		Mode: packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
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
