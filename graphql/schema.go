package graphql

import (
    "github.com/graphql-go/graphql"
)

var userType = graphql.NewObject(graphql.ObjectConfig{
    Name: "User",
    Fields: graphql.Fields{
        "id": &graphql.Field{
            Type: graphql.String,
        },
        "email": &graphql.Field{
            Type: graphql.String,
        },
    },
})

var registerUser = &graphql.Field{
    Type: userType,
    Args: graphql.FieldConfigArgument{
        "email": &graphql.ArgumentConfig{
            Type: graphql.NewNonNull(graphql.String),
        },
        "password": &graphql.ArgumentConfig{
            Type: graphql.NewNonNull(graphql.String),
        },
    },
    Resolve: registerUserResolver,
}

