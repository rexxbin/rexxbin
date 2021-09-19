package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"io/ioutil"
	"log"
	"net/http"
	"rexxbin/database"
	"rexxbin/models"
	"time"
)

func PasteHandler(ctx *gin.Context) {
	pasteID := xid.NewWithTime(time.Now()).String()
	pasteContent, _ := ioutil.ReadAll(ctx.Request.Body)

	pasteModel := models.Paste{
		ID:        pasteID,
		Content:   string(pasteContent),
		CreatedAt: time.Now(),
	}

	_, err := database.MongoInstance.Database("pastes").Collection("data").InsertOne(ctx, pasteModel)
	if err != nil {
		log.Fatalln(err)
	} else {
		ctx.Redirect(http.StatusMovedPermanently, "/v/"+pasteID)
	}
}
