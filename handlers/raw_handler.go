package handlers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"rexxbin/database"
	"rexxbin/models"
)

func RawHandler(ctx *gin.Context) {
	var paste models.Paste
	pasteID := ctx.Param("id")

	db := database.MongoInstance.Database("pastes").Collection("data")
	err := db.FindOne(ctx, bson.M{"_id": pasteID}).Decode(&paste)
	if err != nil {
		log.Fatalln(err)
	}

	ctx.String(http.StatusOK,  paste.Content)
}
