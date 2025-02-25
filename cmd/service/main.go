package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"music/config"
	"music/models"
	"music/utils"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

// Initialize MongoDB
func init() {
	config.InitMongoDB()
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if !utils.IsValidEmail(user.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	if !utils.IsValidPassword(user.Password) {
		http.Error(w, "Password must be at least 8 characters long, contain an uppercase letter, a lowercase letter, and a digit", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user.Password = hashedPassword

	// Store the user in the MongoDB
	collection := config.Client.Database("myapp").Collection("users")
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Fetch the user from the MongoDB
	collection := config.Client.Database("myapp").Collection("users")
	var storedUser models.User
	err = collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&storedUser)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Check password
	err = utils.ComparePassword(storedUser.Password, user.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT
	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Send JWT token as response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func main() {
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/register", registerHandler).Methods("POST")
	r.HandleFunc("/login", loginHandler).Methods("POST")

	// Start the server
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
