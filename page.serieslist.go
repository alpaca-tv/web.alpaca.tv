package main

import "github.com/gin-gonic/gin"

func SerieslistPageHandler(g *gin.Context) {
	g.HTML(200, "page.serieslist.html", nil)
}
