package main

import (
	"log"
	"strconv"

	"github.com/alpaca-tv/alpclib"
	"github.com/gin-gonic/gin"
)

type FilmsPageData struct {
	Order    string
	Films    ComponentBlockCards
	PrevPage int
	NextPage int
}

func FilmsPageHandler(g *gin.Context) {
	// Parse parameters
	order := g.Query("o")
	if order == "" {
		order = "watching"
	}
	pagestr := g.Param("page")
	page, err := strconv.Atoi(pagestr)
	if err != nil {
		log.Panicln("Error while converting page type")
	}
	// Get films
	source := alpclib.Rezka{}
	films, _ := source.ListFilms(&alpclib.ListParameters{
		OrderBy: order,
		Page:    page,
	})
	// Create card components
	filmcards := []ComponentCard{}
	for _, film := range films {
		filmcards = append(filmcards, ComponentCard{
			URL:    "/films/f/" + film.ID,
			Poster: film.PosterURL,
		})
	}
	// Render
	g.HTML(200, "page.films.html", FilmsPageData{
		Order: order,
		Films: ComponentBlockCards{
			Cards: filmcards,
			Wrap:  true,
		},
		NextPage: page + 1,
		PrevPage: page - 1,
	})
}
