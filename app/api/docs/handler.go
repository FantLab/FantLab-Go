package docs

import (
	"fantlab/api/internal/routing"
	"fantlab/pb"
	"go/ast"
	"reflect"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"

	"golang.org/x/tools/go/packages"
)

func getEndpointsInfo(endpoints []routing.Endpoint) (result []string) {
	endpointsPackage := loadPackage("fantlab/api/internal/endpoints")
	pbPackage := loadPackage("fantlab/pb")

	modelComments := makeModelCommentsTable(pbPackage, func(f *ast.Field) bool {
		return !strings.HasPrefix(f.Names[0].Name, "XXX")
	})

	for _, endpoint := range endpoints {
		frame := getCallerFrame(endpoint.Handler())
		if frame == nil {
			continue
		}

		for _, f := range endpointsPackage.Syntax {
			ast.Inspect(f, func(node ast.Node) bool {
				if funcDecl, ok := node.(*ast.FuncDecl); ok {
					if strings.Contains(frame.Function, funcDecl.Name.Name) {
						comment := getFuncComment(f, endpointsPackage.Fset, frame.Line)
						scheme := getEndpointResponseScheme(funcDecl, endpointsPackage, modelComments)
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

func getEndpointResponseScheme(funcDecl *ast.FuncDecl, pkg *packages.Package, modelComments commentsTable) (result string) {
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
				sb := &schemeBuilder{
					indent: "  ",
					getComment: func(t reflect.Type, fieldName string) string {
						typeName := t.String()
						dotIndex := strings.LastIndex(typeName, ".") + 1
						if dotIndex > 0 {
							typeName = typeName[dotIndex:]
						}
						return modelComments[typeName][fieldName]
					},
					isValidField: func(f reflect.StructField) bool {
						return !strings.HasPrefix(f.Name, "XXX")
					},
				}

				result = sb.make(messageType)
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
