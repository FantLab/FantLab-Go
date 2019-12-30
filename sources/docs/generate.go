package docs

import (
	"fantlab/base/routing"
	"fantlab/docs/analysis"
	"io"
	"strconv"
	"strings"
	"text/template"
)

func Generate(w io.Writer, routes *routing.Group, basePath string) error {
	t := template.Must(template.New("markdown").Parse(markdownTemplate))
	data := getTemplateDataFromRoutes(routes, basePath)
	return t.Execute(w, data)
}

func getTemplateDataFromRoutes(routes *routing.Group, basePath string) []t_group {
	var groups []t_group

	routes.Walk(func(group *routing.Group) {
		endpoints := group.Endpoints()
		if len(endpoints) == 0 {
			return
		}

		g := t_group{Info: group.Info()}

		infoTable := analysis.AnalyzeEndpoints(endpoints, "```\n", "\n```")

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
				Path:        basePath + endpoint.Path(),
				File:        relativePathToFile(info.FilePath, info.Line),
				Description: info.Description,
				Params:      params,
				Scheme:      info.ResponseModelScheme,
			}

			g.Endpoints = append(g.Endpoints, e)
		}

		groups = append(groups, g)
	})

	return groups
}

func relativePathToFile(filePath string, line int) string {
	return "../" + filePath[strings.Index(filePath, "sources"):] + "#L" + strconv.Itoa(line)
}
