package helpers

import (
	"bytes"
	"text/template"
)

func InflateTextTemplate(t *template.Template, data interface{}) (string, error) {
	buf := new(bytes.Buffer)
	err := t.Execute(buf, data)

	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
