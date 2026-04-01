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
	"go-graphql-taskflow/internal/loader"
	"go-graphql-taskflow/internal/repository"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	commentRepo := repository.NewCommentRepository()
	commentService := service.NewCommentService(commentRepo)

	taskRepo := repository.NewTaskRepository()
	taskService := service.NewTaskService(taskRepo)
	loaders := &loader.Loaders{
		CommentLoader: loader.NewCommentLoader(commentService),
	}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					CommentService: commentService,
					TaskService: taskService,
				},
			},
		),
	)

	http.Handle(
		"/query",
		DataLoaderMiddleware(loaders, srv),
	)

	// Playground UI
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	// GraphQL endpoint
	
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

func DataLoaderMiddleware(loaders *loader.Loaders, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := loader.InjectLoaders(r.Context(), loaders)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
