package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"

	"github.com/discordianfish/imagen"
	"github.com/discordianfish/imagen/templates"
	"gopkg.in/yaml.v2"
)

var (
	templateDir = flag.String("t", "templates/", "Template directory")

	sanitizeRegex = regexp.MustCompile("[^a-zA-Z0-9_\\-]")
)

func sanitize(str string) string {
	return sanitizeRegex.ReplaceAllString(str, ".")
}

func main() {
	flag.Parse()
	if len(os.Args) < 3 {
		log.Fatal("Usage:", os.Args[0], "path/to/data.yaml")
	}
	configFile := os.Args[1]
	destDir := os.Args[2]
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	conf := &imagen.ConfigFile{}
	if err := yaml.Unmarshal(data, conf); err != nil {
		log.Fatal(err)
	}
	for _, config := range conf.Configs {
		log.Printf("%#v", config)
		filename := path.Join(*templateDir, config.Template, "Dockerfile.tmpl")
		log.Println(filename)
		template, err := templates.New(filename)
		if err != nil {
			log.Fatal(err)
		}
		data := &templates.Data{Labels: &config.Labels}
		for _, base := range config.Bases {
			data.Base.Name = base.Name
			for _, bref := range base.Refs {
				data.Base.Ref = bref
				for _, source := range config.Sources {
					data.Source.Name = source.Name
					for _, ref := range source.Refs {
						data.Source.Ref = ref
						fname := path.Join(
							destDir,
							config.Template,
							sanitize(data.Base.Name),
							sanitize(data.Base.Ref),
							sanitize(data.Source.Name),
							sanitize(data.Source.Ref))
						log.Println("filename:", fname)
						if err := os.MkdirAll(fname, 0755); err != nil {
							log.Fatal(err)
						}
						fh, err := os.Create(path.Join(fname, "Dockerfile"))
						if err != nil {
							log.Fatal(err)
						}
						defer fh.Close()
						if err := template.Execute(fh, data); err != nil {
							log.Fatal(err)
						}
					}
				}
			}
		}
	}
}
