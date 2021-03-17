package main

import (
	"github.com/alpaca-tv/alpclib"
	"github.com/gin-gonic/gin"
)

type IndexPageData struct {
	NWFilms  ComponentBlockCards
	NWSeries ComponentBlockCards
}

func IndexPageHandler(g *gin.Context) {
	r := alpclib.Rezka{}
	nwfilms, _ := r.ListFilms(&alpclib.ListParameters{
		OrderBy: "watching",
	})
	nwserieslist, _ := r.ListSeries(&alpclib.ListParameters{
		OrderBy: "watching",
	})
	nwfilmcards := []ComponentCard{}
	nwserieslistcards := []ComponentCard{}
	for _, film := range nwfilms {
		nwfilmcards = append(nwfilmcards, ComponentCard{
			URL:    "/films/f/" + film.ID,
			Poster: film.PosterURL,
		})
	}
	for _, series := range nwserieslist {
		nwserieslistcards = append(nwserieslistcards, ComponentCard{
			URL:    "/series/s/" + series.ID,
			Poster: series.PosterURL,
		})
	}
	g.HTML(200, "page.index.html", IndexPageData{
		NWFilms: ComponentBlockCards{
			Cards:     nwfilmcards,
			Title:     "Популярные фильмы",
			TitleLink: "/films/1?o=watching",
			EmptyText: "По этому запросу нет результатов",
		},
		NWSeries: ComponentBlockCards{
			Cards:     nwserieslistcards,
			Title:     "Популярные сериалы",
			TitleLink: "/series/1?o=watching",
			EmptyText: "По этому запросу нет результатов",
		},
	})
}
