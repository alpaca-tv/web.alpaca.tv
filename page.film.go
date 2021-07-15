package main

import (
	"github.com/alpaca-tv/alpclib"
	"github.com/gin-gonic/gin"
)

type FilmPageData struct {
	Film alpclib.Film
}

func FilmPageHandler(g *gin.Context) {
	// Parse parameters
	fid := g.Param("id")
	// Get film
	source := alpclib.Rezka{}
	film, _ := source.GetFilm(fid)
	// Render
	g.HTML(200, "page.film.html", &FilmPageData{
		Film: film,
	})
}
