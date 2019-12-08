package docs

type t_endpoint_param struct {
	Name        string
	Source      string
	TypeName    string
	Description string
}

type t_endpoint struct {
	Summary     string
	Method      string
	Path        string
	File        string
	Description string
	Params      []t_endpoint_param
	Scheme      string
}

type t_group struct {
	Info      string
	Endpoints []t_endpoint
}

const markdownTemplate = `
# Список методов

{{range .}}
## {{.Info}}

{{range .Endpoints}}
<details><summary>{{.Summary}}</summary>
<p>

{{.Description}}

**{{.Method}}** [{{.Path}}]({{.File}})
{{if .Params}}
Параметры запроса:

{{range .Params}}
* **{{.Name}}** ({{.Source}}, {{.TypeName}}) - {{.Description}}
{{end}}
{{end}}

Схема ответа:

{{.Scheme}}
---

</p>
</details>
{{end}}
{{end}}
`
