package main

import (
	"strconv"
	"strings"

	"github.com/alpaca-tv/alpclib"
	"github.com/gin-gonic/gin"
)

type FilmViewPageData struct {
	Film           alpclib.Film
	SourceSelector ComponentSelectorSourceFilm
}

func FilmSourceIndexByPriority(sources []alpclib.FilmSource) int {
	// Variant indexes
	sindexes := []int{}
	for i := 0; i < len(sources); i++ {
		sindexes = append(sindexes, i)
	}
	// Prioritize dubbing
	vcovers := []string{"Дубляж (реж. версия)", "Дубляж"}
	for _, vc := range vcovers {
		_sindexes := []int{}
		for _, si := range sindexes {
			if strings.Contains(sources[si].Voicecover, vc) {
				_sindexes = append(_sindexes, si)
			}
		}
		if len(_sindexes) != 0 {
			sindexes = _sindexes
		}
	}
	// Prioritize quality
	qualities := []string{"1080", "720", "480", "360", "240"}
	for _, q := range qualities {
		_sindexes := []int{}
		for _, si := range sindexes {
			if strings.Contains(sources[si].Quality, q) {
				_sindexes = append(_sindexes, si)
			}
		}
		if len(_sindexes) != 0 {
			sindexes = _sindexes
		}
	}
	// Return index
	return sindexes[0]
}

func FilmViewPageHandler(g *gin.Context) {
	// Parse parameters
	fid := g.Param("id")
	// Get film
	source := &alpclib.Rezka{}
	film, _ := source.GetFilm(fid)
	// Determine source index
	srcindex := 0
	if g.Query("source") != "" {
		srcindex, _ = strconv.Atoi(g.Query("source"))
	} else {
		srcindex = FilmSourceIndexByPriority(film.Sources)
	}
	// Render
	g.HTML(200, "page.film.view.html", &FilmViewPageData{
		Film: film,
		SourceSelector: ComponentSelectorSourceFilm{
			Sources:  film.Sources,
			Selected: srcindex,
		},
	})
}
