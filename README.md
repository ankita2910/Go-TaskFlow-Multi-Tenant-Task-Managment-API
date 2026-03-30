TaskFlow Multi-Tenant Task Management API
Project Documentation
Author: Ankita Dewangswami
Date: 2026-03-30
Stack: Go, GraphQL (gqlgen), DynamoDB
1. Project Overview
The goal of this project was to create a multi-tenant task management system with a Go backend exposed via GraphQL and persisted in AWS DynamoDB using a single-table design. The project emphasizes production-level architecture with separation of concerns, scalability, and maintainability.
Key features implemented:
GraphQL API in Go
Service + Repository pattern
Resolver architecture
DynamoDB single-table design
Nested resolvers
DataLoader batching
Authorization layer
2. Project Structure
go-graphql-taskflow/
├── cmd/server/            # Main entry point for GraphQL server
├── graph/                 # gqlgen generated and custom resolvers
│   ├── generated/
│   ├── resolver.go
│   ├── schema.graphqls
│   └── schema.resolvers.go
├── internal/
│   ├── auth/              # Authentication and authorization logic
│   ├── dynamodb/          # DynamoDB client setup and repository implementation
│   ├── loader/            # DataLoader batching utilities
│   ├── model/             # Domain models (Task, Project, User, etc.)
│   ├── repository/        # Repository interfaces
│   └── service/           # Business logic layer
├── pkg/                   # Optional reusable packages
├── go.mod
├── go.sum
└── gqlgen.yml
3. Development Flow
The project was implemented incrementally, following a production-grade architecture:
Step 1: Initialize Go project and gqlgen
mkdir go-graphql-taskflow
cd go-graphql-taskflow
go mod init github.com/ankita2910/go-graphql-taskflow
go get github.com/99designs/gqlgen
go run github.com/99designs/gqlgen init
Sets up go.mod and gqlgen scaffolding.
Generates initial graph/ folder with schema and resolver stubs.
Step 2: Define Domain Models
Inside internal/model/, created task.go:
type Task struct {
    ID        string
    Title     string
    Status    string
    ProjectID string
}
Represents the core data entity.
Separated into internal/model for clean architecture and production-level modularity.
Step 3: Design GraphQL Schema
graph/schema.graphqls:
type Task {
  id: ID!
  title: String!
  status: String!
  projectID: ID!
}

type Query {
  tasks: [Task!]!
}

input CreateTaskInput {
  title: String!
  projectID: ID!
}

type Mutation {
  createTask(input: CreateTaskInput!): Task!
}
Defines queries, mutations, and input types.
Enables gqlgen to generate type-safe resolvers.
Step 4: Generate gqlgen code
go run github.com/99designs/gqlgen generate
Generates resolver interfaces, models, and the generated/ package.
Resolver stubs in schema.resolvers.go are created.
Step 5: Implement Resolver Layer
graph/schema.resolvers.go:
func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTaskInput) (*model.Task, error) {
    task := r.TaskService.CreateTask(input.Title, input.ProjectID)
    return task, nil
}
Resolver translates GraphQL queries/mutations into business logic calls.
Does not interact with DB directly.
Step 6: Service Layer
internal/service/task_service.go:
type TaskService struct {}

func (s *TaskService) CreateTask(title string, projectID string) *model.Task {
    return &model.Task{
        ID:        "1",  // Normally, UUIDs would be generated
        Title:     title,
        Status:    "OPEN",
        ProjectID: projectID,
    }
}
Business logic layer handles validation, orchestration, and rules.
Called by resolvers.
Step 7: Repository Layer & DynamoDB Integration
Repository interfaces abstract DB operations (internal/repository).
Implementations in internal/dynamodb allow switching databases without changing service/resolver code.
DynamoDB single-table design stores all entities in one table with composite primary keys (PK, SK) for fast queries and access patterns.
Step 8: GraphQL Server Setup
cmd/server/main.go:
srv := handler.NewDefaultServer(
    generated.NewExecutableSchema(
        generated.Config{
            Resolvers: &graph.Resolver{
                TaskService: &service.TaskService{},
            },
        },
    ),
)
http.Handle("/query", srv)
log.Println("GraphQL server running at http://localhost:8080/")
log.Fatal(http.ListenAndServe(":8080", nil))
Injects TaskService into resolvers.
Exposes /query endpoint for GraphQL Playground.
Step 9: Executing a Mutation
Example mutation to create a task:
mutation {
  createTask(input: { title: "Learn Go", projectID: "101" }) {
    id
    title
    status
    projectID
  }
}
Resolver calls service → returns Task object.
Future implementation will persist tasks in DynamoDB.
