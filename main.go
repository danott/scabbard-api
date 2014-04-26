package main

import (
	"flag"
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/danott/envflag"
)

var apiKey = "IP"

func main() {
	flag.StringVar(&apiKey, "api-key", "IP", "API Key for the ESV API")
	envflag.Parse()

	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/search", PassageQueryHandler)
	m.Get("/", HelpHandler)
	m.Run()
}

type arbitraryJSON map[string]interface{}

func PassageQueryHandler(ren render.Render, req *http.Request) {
	passageQuery, err := PassageQuery(req.URL.Query().Get("q"))

	if err != nil {
		ren.JSON(404, arbitraryJSON{"status": 404, "message": err.Error()})
	} else {
		ren.JSON(200, passageQuery)
	}
}

func HelpHandler(ren render.Render) {
	ren.JSON(200, arbitraryJSON{"help": "Use endpoint '/search?q=query' to search for a passage."})
}
