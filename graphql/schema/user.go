package schema

import (
	"context"
	"errors"
	"music/config"
	"music/models"
	"music/utils"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
)

// Define the User Type
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// Define the input type for registration and login
var UserInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"password": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

// Define the Mutation for Registration
var registerMutation = &graphql.Field{
	Type: UserType,
	Args: graphql.FieldConfigArgument{
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		email := p.Args["email"].(string)
		password := p.Args["password"].(string)

		// Validate email and password
		if !utils.IsValidEmail(email) {
			return nil, errors.New("Invalid email format")
		}

		if !utils.IsValidPassword(password) {
			return nil, errors.New("Password must be at least 8 characters long and contain uppercase, lowercase, and a digit")
		}

		// Hash the password
		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			return nil, err
		}

		// Store user in MongoDB
		user := models.User{
			Email:    email,
			Password: hashedPassword,
		}

		collection := config.Client.Database("myapp").Collection("users")
		_, err = collection.InsertOne(context.Background(), user)
		if err != nil {
			return nil, err
		}

		return &user, nil
	},
}

// Define the Mutation for Login
var loginMutation = &graphql.Field{
	Type: graphql.String, // Returns JWT Token
	Args: graphql.FieldConfigArgument{
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		email := p.Args["email"].(string)
		password := p.Args["password"].(string)

		// Find user by email in MongoDB
		collection := config.Client.Database("myapp").Collection("users")
		var user models.User
		err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
		if err != nil {
			return nil, errors.New("User not found")
		}

		// Compare password with stored hashed password
		err = utils.ComparePassword(user.Password, password)
		if err != nil {
			return nil, errors.New("Invalid credentials")
		}

		// Generate JWT token
		token, err := utils.GenerateJWT(email)
		if err != nil {
			return nil, err
		}

		return token, nil
	},
}

// Define the root mutation
var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"register": registerMutation,
		"login":    loginMutation,
	},
})

// Define the Root Schema
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    nil, // No queries here, but you can add user queries if needed
	Mutation: Mutation,
})
