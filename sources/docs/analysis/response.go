package analysis

import (
	"fantlab/pb"
	"go/ast"
	"reflect"
	"strings"

	"golang.org/x/tools/go/packages"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func getResponseModelType(funcDecl *ast.FuncDecl, pkg *packages.Package) reflect.Type {
	var responseType reflect.Type

	ast.Inspect(funcDecl, func(node ast.Node) bool {
		if rtrn, ok := node.(*ast.ReturnStmt); ok {
			var messageType reflect.Type

			for i, expr := range rtrn.Results {
				if t, ok := pkg.TypesInfo.Types[expr]; ok {
					switch i {
					case 0:
						continue
					case 1:
						if t.Type != nil {
							protoName := getProtoNameFromTypeName(t.Type.String())

							if mt, _ := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(protoName)); mt != nil {
								messageType = reflect.TypeOf(mt.Zero().Interface())
							}
						}
					default:
						break
					}
				}
			}

			if messageType != nil && messageType != reflect.TypeOf((*pb.Error_Response)(nil)) {
				responseType = messageType
			}
		}
		return responseType == nil
	})

	return responseType
}

func getProtoNameFromTypeName(typeName string) string {
	typeName = typeName[(strings.Index(typeName, ".") + 1):]
	return strings.ReplaceAll(typeName, "_", ".")
}
