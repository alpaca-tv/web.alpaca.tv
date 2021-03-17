package main

import (
	"log"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// Init Gin
	g := gin.New()
	// Register services
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	// Register templates
	g.LoadHTMLGlob("*.html")
	// Register statics
	g.Use(static.Serve("/static/", static.LocalFile("./static", true)))
	// Register pages
	g.GET("/", IndexPageHandler)
	g.GET("/search", SearchPageHandler)
	g.GET("/films/:page", FilmsPageHandler)
	g.GET("/films/:page/:id", FilmPageHandler)
	g.GET("/series/:page", SerieslistPageHandler)
	g.GET("/series/:page/:id", SeriesPageHandler)
	// Run
	if err := g.Run(":25025"); err != nil {
		log.Panicln(err)
	}
}
