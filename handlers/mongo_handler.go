package handlers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

const connectionString = ""

var mongoInstance sync.Once
var clientInstance *mongo.Client
var clientError error

func GetMongoClient() {
	mongoInstance.Do(func() {
		clientOptions := options.Client().ApplyURI(connectionString)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientError = err
		}

		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientError = err
		}
		clientInstance = client
	})
}
