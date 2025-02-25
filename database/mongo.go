package database

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Connect establishes a connection to MongoDB
func Connect() (*mongo.Client, error) {
    uri := "mongodb://localhost:27017"
    clientOptions := options.Client().ApplyURI(uri)
    var err error
    client, err = mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        return nil, err
    }

    err = client.Ping(context.Background(), nil)
    if err != nil {
        return nil, err
    }

    fmt.Println("Connected to MongoDB!")
    return client, nil
}

// GetCollection returns a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
    return client.Database("mydb").Collection(collectionName)
}
