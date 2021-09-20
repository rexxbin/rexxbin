package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"rexxbin/handlers"
)

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8000"
}

func main() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.Use(gin.Logger())

	router.LoadHTMLGlob("static/*.html")
	router.Static("/css", "static/css/")
	router.Static("/js", "static/js/")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Rexxbin - Pastebin written in Go",
			"year":  2021,
		})
	})

	router.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "Api is Alive and Running",
		})
	})

	router.POST("/api/paste", handlers.PasteHandler)

	router.GET("/v/:id", handlers.ViewHandler)

	router.GET("/v/:id/raw", func(ctx *gin.Context) {
	})

	err := router.Run(getPort())
	if err != nil {
		log.Fatalln(err)
	}
}
