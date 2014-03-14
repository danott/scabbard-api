package main

import (
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/search", PassageQueryHandler)
	m.Run()
}

func PassageQueryHandler(ren render.Render, req *http.Request) {
	ren.JSON(200, PassageQuery(req.URL.Query().Get("q")))
}
