package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rexxbin/handlers"
)

func main() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.LoadHTMLGlob("static/*.html")
	router.Static("/css", "static/css/")
	router.Static("/js", "static/js/")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Rexxbin - Pastebin Service",
			"year":  2021,
		})
	})

	router.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "Api is Alive and Running",
		})
	})

	router.POST("/api/paste", handlers.PasteHandler)

	router.GET("/:id", handlers.ViewHandler)

	err := router.Run(":8000")
	if err != nil {
		log.Fatalln(err)
	}
}
