package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"go-graphql-taskflow/internal/graph"
	"go-graphql-taskflow/internal/graph/generated"
	"go-graphql-taskflow/internal/service"
	"fmt"
	"go-graphql-taskflow/internal/graph/model"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Create GraphQL server
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					TaskService: &service.TaskService{},
				},
			},
		),
	)

	// Playground UI
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	// GraphQL endpoint
	http.Handle("/query", srv)
    task := model.Task{
        ID:        "1",
        Title:     "Learn Go",
        Status:    "OPEN",
        ProjectID: "101",
    }
    fmt.Println(task)
	log.Printf("GraphQL server running at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}