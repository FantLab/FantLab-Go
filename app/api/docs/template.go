package docs

type t_endpoint struct {
	Summary string
	Method  string
	Path    string
	File    string
	Text    string
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

**{{.Method}}** [{{.Path}}]({{.File}})

{{.Text}}

---

</p>
</details>
{{end}}
{{end}}
`
