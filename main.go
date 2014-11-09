package main

import (
	"flag"
	"net/http"
	"net/http/fcgi"
	"runtime"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/danott/envflag"
)

var apiKey = "IP"

type queryResult struct {
	passage Passage
	err     error
}

type simpleStore map[string]queryResult

func main() {
	flag.StringVar(&apiKey, "api-key", "IP", "API Key for the ESV API")
	envflag.Parse()

	m := martini.Classic()
	m.Use(render.Renderer())

	store := make(simpleStore)
	m.Map(store)

	m.Get("/search", PassageQueryHandler)
	m.Get("/", HelpHandler)

	if runtime.GOOS == "linux" {
		fcgi.Serve(nil, m)
	} else {
		m.Run()
	}
}

type arbitraryJSON map[string]interface{}

func PassageQueryHandler(ren render.Render, req *http.Request, store simpleStore) {
	var passage Passage
	var err error
	q := req.URL.Query().Get("q")

	if v, ok := store[q]; ok {
		passage, err = v.passage, v.err
	} else {
		passage, err = PassageQuery(q)
		store[q] = queryResult{passage, err}
	}

	if err != nil {
		ren.JSON(404, arbitraryJSON{"status": 404, "message": err.Error()})
	} else {
		ren.JSON(200, passage)
	}
}

func HelpHandler(ren render.Render) {
	ren.JSON(200, arbitraryJSON{"help": "Use endpoint '/search?q=query' to search for a passage."})
}
