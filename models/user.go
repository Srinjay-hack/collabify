package models

import (
    "context"
    "music/database"
    "go.mongodb.org/mongo-driver/bson"
)

type User struct {
    ID       string `json:"id" bson:"_id,omitempty"`
    Email    string `json:"email" bson:"email"`
    Password string `json:"password" bson:"password"`
}

// CreateUser adds a new user to the MongoDB collection
func CreateUser(user User) error {
    collection := database.GetCollection("users")
    _, err := collection.InsertOne(context.Background(), user)
    return err
}

// GetUserByEmail retrieves a user by email
func GetUserByEmail(email string) (User, error) {
    collection := database.GetCollection("users")
    var user User
    err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
    return user, err
}
