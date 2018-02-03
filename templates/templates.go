package templates

import (
	"io"
	"path"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"escapeQuote": func(str string) string {
		return strings.Replace(str, "\"", "\\\"", -1)
	},
}

type Template struct {
	name string
	*template.Template
}

func New(filename string) (*Template, error) {
	name := path.Base(filename)
	t, err := template.New(name).Funcs(funcMap).ParseFiles(filename)
	if err != nil {
		return nil, err
	}
	return &Template{name, t}, nil
}

// FIXME: Use interface type for data
func (t *Template) Execute(w io.Writer, data interface{}) error {
	return t.Template.ExecuteTemplate(w, t.name, data)
}
