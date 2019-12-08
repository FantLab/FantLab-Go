package docs

import (
	"fantlab/server/internal/docs/analysis"
	"fantlab/server/internal/routing"
	"io"
	"strconv"
	"strings"
	"text/template"
)

func Generate(w io.Writer) error {
	t := template.Must(template.New("markdown").Parse(markdownTemplate))
	data := getTemplateDataFromRoutes()
	return t.Execute(w, data)
}

func getTemplateDataFromRoutes() []t_group {
	var groups []t_group

	routing.Routes(nil, nil, nil).Walk(func(group *routing.Group) {
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
				Path:        routing.BasePath + endpoint.Path(),
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
