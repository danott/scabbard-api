package main

import (
	"flag"
	"log"
	"net/http"

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

// DB Returns a martini.Handler
func DB() martini.Handler {
	db := make(simpleStore)

	return func(c martini.Context) {
		c.Map(db)
		c.Next()
	}
}
func main() {
	flag.StringVar(&apiKey, "api-key", "IP", "API Key for the ESV API")
	envflag.Parse()

	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(DB())

	m.Get("/search", PassageQueryHandler)
	m.Get("/", HelpHandler)
	m.Run()
}

type arbitraryJSON map[string]interface{}

func PassageQueryHandler(ren render.Render, req *http.Request, store simpleStore) {
	var passage Passage
	var err error

	q := req.URL.Query().Get("q")
	if v, ok := store[q]; ok {
		log.Printf("Found in cache")
		passage, err = v.passage, v.err
	} else {
		log.Printf("Found in api")
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
