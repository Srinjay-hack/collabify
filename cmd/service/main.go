package service

import (
    "fmt"
    "log"
    "music/handlers"
    "music/database"
)

func main() {
    // Connect to MongoDB
    client, err := database.Connect()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB!")

    // Start GraphQL handler
    handlers.StartGraphQLServer(client)
}
