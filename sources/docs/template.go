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

type t_enum_case struct {
	Num  int
	Name string
}

type t_enum struct {
	Name  string
	Cases []t_enum_case
}

type t_content struct {
	Enums  []t_enum
	Groups []t_group
}

const markdownTemplate = `
# Константы

{{range .Enums}}
<details><summary>{{.Name}}</summary>
<p>

| Int | String |
| --- | --- |{{range .Cases}}
| {{.Num}} | {{.Name}} |{{end}}
---

</p>
</details>
{{end}}

# Эндпойнты

{{range .Groups}}
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
