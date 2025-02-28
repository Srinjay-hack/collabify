package config

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joho/godotenv"
)

var Client *mongo.Client

// Initialize MongoDB connection
func InitMongoDB() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	err = godotenv.Load(filepath.Join(pwd, "../../.env"))
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbURI := os.Getenv("DATABASE_URI")

	//uri := "mongodb://localhost:27017/collabify" // Change this to your MongoDB URI
	client, err := mongo.NewClient(options.Client().ApplyURI(dbURI))
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
