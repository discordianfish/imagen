package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/julienschmidt/httprouter"

	"github.com/discordianfish/imagen"
	"github.com/discordianfish/imagen/templates"
)

var (
	templateDir = flag.String("d", "templates/", "Template directory")
	listenAddr  = flag.String("l", ":8080", "Address to listen on")
)

func main() {
	flag.Parse()
	router := httprouter.New()
	// router.GET("/", handleIndex)

	dh, err := os.Open(*templateDir)
	if err != nil {
		log.Fatal(err)
	}
	dirs, err := dh.Readdir(-1)
	dh.Close()
	if err != nil {
		log.Fatal(err)
	}
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		fname := dir.Name()
		filename := path.Join(*templateDir, fname, "Dockerfile.tmpl")
		t, err := templates.New(filename)
		if err != nil {
			log.Fatal(err)
		}
		router.GET(path.Join("/", fname, ":ref/*origin"), handler(t))
	}
	log.Fatal(http.ListenAndServe(*listenAddr, router))
}

func handler(t *templates.Template) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		data := &imagen.Config{
			Base: imagen.Base{Version: "1.9"},
			Source: imagen.Source{
				Origin: ps.ByName("origin")[1:], //FIXME: Why?
				Ref:    ps.ByName("ref"),
			},
		}

		// FIXME: We probably need to buffer this, so we don't return
		// 2xx implicitly and then run into an error.
		if err := t.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
