package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rektbin/handlers"
)

func main() {
	handlers.GetMongoClient()
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.LoadHTMLGlob("static/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "BinGO - Pastebin Service",
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
