package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// Connect to MongoDB
func ConnectDB() {
	var err error
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	log.Println("Successfully connected to MongoDB")
}
