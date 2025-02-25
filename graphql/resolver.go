// Resolver for user registration
package graphql

import (
	"music/models"
	"music/utils"

	"github.com/graphql-go/graphql"
)

func registerUserResolver(params graphql.ResolveParams) (interface{}, error) {
	email := params.Args["email"].(string)
	password := params.Args["password"].(string)

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Email:    email,
		Password: hashedPassword,
	}

	err = models.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name:   "Mutation",
	Fields: graphql.Fields{"registerUser": registerUser},
})

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:   "RootQuery",
	Fields: graphql.Fields{},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    RootQuery,
	Mutation: Mutation,
})
