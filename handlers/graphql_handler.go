package handlers

import (
	"log"
	"music/graphql"
	"net/http"

	"github.com/graphql-go/handler"
	"go.mongodb.org/mongo-driver/mongo"
)

// StartGraphQLServer starts the GraphQL server on port 8080
func StartGraphQLServer(client *mongo.Client) {
    graphQLHandler := handler.New(&handler.Config{
        Schema: &graphql.Schema,
        Pretty: true,
    })

    http.Handle("/graphql", graphQLHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
