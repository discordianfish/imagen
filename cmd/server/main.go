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

	dh, err := os.Open(*templateDir)
	if err != nil {
		log.Fatal(err)
	}
	dirs, err := dh.Readdir(-1)
	dh.Close()
	if err != nil {
		log.Fatal(err)
	}
	routes := []string{}
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
		p := path.Join("/", fname, ":ref/*name")
		routes = append(routes, p)
		router.GET(p, handler(t))
	}
	filename := path.Join(*templateDir, "index.tmpl")
	t, err := templates.New(filename)
	if err != nil {
		log.Fatal(err)
	}
	router.GET("/", indexHandler(t, routes))
	log.Fatal(http.ListenAndServe(*listenAddr, router))
}

func indexHandler(t *templates.Template, routes []string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if err := t.Execute(w, routes); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func handler(t *templates.Template) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		data := &imagen.Config{
			Source: imagen.Source{
				Name: ps.ByName("name")[1:], //FIXME: Why?
				Ref:  ps.ByName("ref"),
			},
		}

		// FIXME: We probably need to buffer this, so we don't return
		// 2xx implicitly and then run into an error.
		if err := t.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
