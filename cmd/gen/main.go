package main

import (
	"log"
	"os"

	"github.com/discordianfish/imagen"
	"github.com/discordianfish/imagen/templates"
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

	data := &imagen.Config{}
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
