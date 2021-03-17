package main

import "github.com/gin-gonic/gin"

func FilmPageHandler(g *gin.Context) {
	g.HTML(200, "page.film.html", nil)
}
