package analysis

import (
	"fantlab/docs/scheme"
	"fantlab/server/routing"
	"fmt"
	"go/ast"
	"reflect"
	"strings"
	"sync"
	"unicode"

	"golang.org/x/tools/go/packages"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ParamInfo struct {
	Name        string
	TypeName    string
	Source      string
	Description string
}

type EndpointInfo struct {
	FilePath            string
	Line                int
	Description         string
	Params              []*ParamInfo
	ResponseModelScheme string
}

var (
	endpointsPackage   *packages.Package
	protoModelsPackage *packages.Package
	loadPackagesOnce   sync.Once
)

func isValidFieldName(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		} else {
			break
		}
	}
	return false
}

func AnalyzeEndpoints(endpoints []routing.Endpoint, schemePrefix, schemePostfix string, enumsTable map[string]protoreflect.EnumValueDescriptors) map[int]*EndpointInfo {
	loadPackagesOnce.Do(func() {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			endpointsPackage = loadPackage("fantlab/server/internal/endpoints")
			wg.Done()
		}()
		go func() {
			protoModelsPackage = loadPackage("fantlab/pb")
			wg.Done()
		}()
		wg.Wait()
	})

	endpointsTable := make(map[int]*EndpointInfo)

	modelComments := makeModelCommentsTable(protoModelsPackage, func(f *ast.Field) bool {
		return isValidFieldName(f.Names[0].Name)
	})
	schemeBuilder := makeSchemeBuilder(modelComments, enumsTable)
	funcDecls := collectFuncDecls(endpointsPackage)

	for index, endpoint := range endpoints {
		frame := getCallerFrame(endpoint.Handler())
		if frame == nil {
			continue
		}

		for funcName, funcDecl := range funcDecls {
			if !strings.Contains(frame.Function, "."+funcName+"-") {
				continue
			}

			responseType := getResponseModelType(funcDecl, endpointsPackage)
			if responseType == nil {
				continue
			}

			params := getRequestParams(funcDecl, endpointsPackage)
			scheme := schemeBuilder.Make(responseType, schemePrefix, schemePostfix)

			endpointsTable[index] = &EndpointInfo{
				FilePath:            frame.File,
				Line:                frame.Line,
				Description:         funcDecl.Doc.Text(),
				Params:              params,
				ResponseModelScheme: scheme,
			}

			break
		}
	}

	return endpointsTable
}

func makeSchemeBuilder(modelComments commentsTable, enumsTable map[string]protoreflect.EnumValueDescriptors) *scheme.Builder {
	return scheme.NewBuilder(&scheme.BuilderConfig{
		GetComment: func(t reflect.Type, fieldName string) string {
			typeName := t.String()
			dotIndex := strings.LastIndex(typeName, ".") + 1
			if dotIndex > 0 {
				typeName = typeName[dotIndex:]
			}
			comment := modelComments[typeName][fieldName]
			if "" != comment {
				return " # " + comment
			}
			return ""
		},
		IsValidField: func(f reflect.StructField) bool {
			return isValidFieldName(f.Name)
		},
		GetFieldName: func(tag reflect.StructTag) string {
			var jsonName string
			for _, s := range strings.Split(tag.Get("protobuf"), ",") {
				if strings.HasPrefix(s, "name") {
					if jsonName == "" {
						jsonName = strings.Split(s, "=")[1]
					}
				} else if strings.HasPrefix(s, "json") {
					jsonName = strings.Split(s, "=")[1]
				}
			}
			return jsonName
		},
		CustomStructFormatter: func(t reflect.Type) string {
			if t.Name() == "Timestamp" {
				return "timestamp"
			}
			return ""
		},
		CustomIntFormatter: func(t reflect.Type) string {
			if method, exists := t.MethodByName("Descriptor"); exists {
				if _, exists = enumsTable[t.Name()]; !exists {
					enumsTable[t.Name()] = method.Func.Call([]reflect.Value{reflect.Indirect(reflect.New(t))})[0].Interface().(protoreflect.EnumDescriptor).Values()
				}
				return fmt.Sprintf("enum (%s)", t.Name())
			}
			return ""
		},
	})
}

func collectFuncDecls(pkg *packages.Package) map[string]*ast.FuncDecl {
	table := make(map[string]*ast.FuncDecl)

	for _, f := range pkg.Syntax {
		ast.Inspect(f, func(node ast.Node) bool {
			if funcDecl, ok := node.(*ast.FuncDecl); ok {
				table[funcDecl.Name.Name] = funcDecl
				return false
			}
			return true
		})
	}

	return table
}
