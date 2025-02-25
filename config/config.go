package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// Initialize MongoDB connection
func InitMongoDB() {
	uri := "mongodb://localhost:27017" // Change this to your MongoDB URI
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Error creating MongoDB client:", err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	Client = client
	fmt.Println("Connected to MongoDB!")
}
