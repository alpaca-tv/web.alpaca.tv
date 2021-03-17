package main

import "github.com/gin-gonic/gin"

func SeriesPageHandler(g *gin.Context) {
	g.HTML(200, "page.series.html", nil)
}
