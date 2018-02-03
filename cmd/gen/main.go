package main

import (
	"log"
	"os"

	"github.com/discordianfish/imagen/templates"
	"github.com/discordianfish/imagen/templates/golang"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage:", os.Args[0], "path/to/template")
	}
	filename := os.Args[1]
	t, err := templates.New(filename)
	if err != nil {
		log.Fatal(err)
	}

	data := &golang.Config{}
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
