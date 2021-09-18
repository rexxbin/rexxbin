package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"log"
	"net/http"
	"rektbin/models"
	"time"
)

func PasteHandler(ctx *gin.Context) {
	pasteID := xid.NewWithTime(time.Now()).String()
	pasteContent := ctx.PostForm("content")

	pasteModel := models.Paste{
		ID:        pasteID,
		Content:   pasteContent,
		CreatedAt: time.Now(),
	}

	_, err := clientInstance.Database("pastes").Collection("data").InsertOne(ctx, pasteModel)
	if err != nil {
		log.Fatalln(err)
	} else {
		ctx.Redirect(http.StatusMovedPermanently, "/"+pasteID)
	}
}
