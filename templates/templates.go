package templates

import (
	"io"
	"path"
	"strings"
	"text/template"

	"github.com/discordianfish/imagen"
)

var funcMap = template.FuncMap{
	"default": func(str, def string) string {
		if str == "" {
			return def
		}
		return str
	},
	"escapeQuote": func(str string) string {
		return strings.Replace(str, "\"", "\\\"", -1)
	},
}

type Data struct {
	Base struct {
		Name string `yaml:"name"`
		Ref  string `yaml:"ref"`
	}
	Source struct {
		Name string `yaml:"name"`
		Ref  string `yaml:"ref"`
	}
	*imagen.Labels
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
