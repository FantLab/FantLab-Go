package docs

import (
	"fantlab/apiserver/routing"
	"fantlab/docs/analysis"
	"io"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func Generate(w io.Writer, routes *routing.Group, basePath string) error {
	t := template.Must(template.New("markdown").Parse(markdownTemplate))
	content := getTemplateDataFromRoutes(routes, basePath)
	return t.Execute(w, content)
}

func getTemplateDataFromRoutes(routes *routing.Group, basePath string) *t_content {
	var groups []t_group

	enumsTable := make(map[string]protoreflect.EnumValueDescriptors)

	routes.Walk(func(group *routing.Group) {
		endpoints := group.Endpoints()
		if len(endpoints) == 0 {
			return
		}

		g := t_group{Info: group.Info()}

		infoTable := analysis.AnalyzeEndpoints(endpoints, "```\n", "\n```", enumsTable)

		for index, endpoint := range endpoints {
			info := infoTable[index]
			if info == nil {
				continue
			}

			var params []t_endpoint_param

			for _, p := range info.Params {
				d := p.Description
				if "" == d {
					d = "нет описания"
				}

				params = append(params, t_endpoint_param{
					Name:        p.Name,
					Source:      p.Source,
					TypeName:    p.TypeName,
					Description: d,
				})
			}

			e := t_endpoint{
				Summary:     endpoint.Info(),
				Method:      endpoint.Method(),
				Path:        basePath + patchPath(endpoint.Path()),
				File:        relativePathToFile(info.FilePath, info.Line),
				Description: info.Description,
				Params:      params,
				Scheme:      info.ResponseModelScheme,
			}

			g.Endpoints = append(g.Endpoints, e)
		}

		groups = append(groups, g)
	})

	var enums []t_enum
	{
		for name, values := range enumsTable {
			var cases []t_enum_case
			for i := 0; i < values.Len(); i++ {
				value := values.Get(i)
				value.Number()
				cases = append(cases, t_enum_case{
					Num:  int(value.Number()),
					Name: string(value.Name()),
				})
			}
			sort.Slice(cases, func(i, j int) bool { return cases[i].Num < cases[j].Num })
			enums = append(enums, t_enum{
				Name:  name,
				Cases: cases,
			})
		}
		sort.Slice(enums, func(i, j int) bool { return strings.Compare(enums[i].Name, enums[j].Name) < 0 })
	}
	return &t_content{
		Enums:  enums,
		Groups: groups,
	}
}

func relativePathToFile(filePath string, line int) string {
	return "../" + filePath[strings.Index(filePath, "sources"):] + "#L" + strconv.Itoa(line)
}

func patchPath(path string) string {
	segments := strings.FieldsFunc(path, func(r rune) bool {
		return r == '/'
	})
	var sb strings.Builder
	sb.Grow(len(path))
	for _, segment := range segments {
		sb.WriteRune('/')
		if segment[0] == ':' {
			sb.WriteRune('{')
			sb.WriteString(segment[1:])
			sb.WriteRune('}')
		} else {
			sb.WriteString(segment)
		}
	}
	return sb.String()
}
