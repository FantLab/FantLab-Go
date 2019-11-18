package docs

import (
	"fantlab/api/internal/routing"
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

func getTemplateDataFromRoutes() (result []t_group) {
	routing.Routes(nil, nil, nil).Walk(func(group *routing.Group) {
		endpoints := group.Endpoints()

		if len(endpoints) == 0 {
			return
		}

		g := t_group{Info: group.Info()}

		for _, endpoint := range endpoints {
			info := getHandlerInfo(endpoint.Handler())

			text := info.comment
			if text == "" {
				text = "Нет описания"
			}

			e := t_endpoint{
				Summary: endpoint.Info(),
				Method:  endpoint.Method(),
				Path:    routing.BasePath + endpoint.Path(),
				File:    relativePathToFile(info.file, info.line),
				Text:    text,
			}

			g.Endpoints = append(g.Endpoints, e)
		}

		result = append(result, g)
	})

	return
}

func relativePathToFile(filePath string, line int) string {
	return "../" + filePath[strings.Index(filePath, "app"):] + "#L" + strconv.Itoa(line)
}
