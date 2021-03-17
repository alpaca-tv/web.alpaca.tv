package main

import (
	"github.com/alpaca-tv/alpclib"
	"github.com/gin-gonic/gin"
)

type SearchPageData struct {
	Films  ComponentBlockCards
	Series ComponentBlockCards
}

func SearchPageHandler(g *gin.Context) {
	source := alpclib.Rezka{}
	films, _ := source.ListFilms(&alpclib.ListParameters{
		Search: g.Query("q"),
	})
	serieslist, _ := source.ListSeries(&alpclib.ListParameters{
		Search: g.Query("q"),
	})
	filmcards := []ComponentCard{}
	serieslistcards := []ComponentCard{}
	for _, film := range films {
		filmcards = append(filmcards, ComponentCard{
			URL:    "/films/f/" + film.ID,
			Poster: film.PosterURL,
		})
	}
	for _, series := range serieslist {
		serieslistcards = append(serieslistcards, ComponentCard{
			URL:    "/series/s/" + series.ID,
			Poster: series.PosterURL,
		})
	}
	g.HTML(200, "page.search.html", SearchPageData{
		Films: ComponentBlockCards{
			Cards:     filmcards,
			Title:     "Фильмы",
			EmptyText: "По этому запросу нет результатов",
			Wrap:      true,
		},
		Series: ComponentBlockCards{
			Cards:     serieslistcards,
			Title:     "Сериалы",
			EmptyText: "По этому запросу нет результатов",
			Wrap:      true,
		},
	})
}
