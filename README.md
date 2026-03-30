Go + GraphQL Task Management API — Architecture and Implementation Document
Project Overview
The Go-GraphQL-TaskFlow project is a multi-layered backend application implemented in Go using GraphQL. The goal of the project is to demonstrate modern backend engineering best practices, including:
Layered architecture (Resolvers → Service → Repository)
Input validation and business logic separation
Clean dependency injection
Preparatory work for database integration (e.g., DynamoDB)
Scalability and maintainability principles
The system currently supports:
Creating a task via a GraphQL mutation
Placeholder query resolver for fetching tasks
1. Project Structure
The project follows an interview-ready, production-oriented structure:
go-graphql-taskflow/
├── cmd/server/             # Application entry point
│   └── main.go
├── internal/
│   ├── auth/               # Authorization logic (future)
│   ├── dynamodb/           # DynamoDB client & repository (future)
│   ├── loader/             # DataLoader for batch queries (future)
│   ├── model/              # Domain models (Task, User, etc.)
│   ├── repository/         # Interfaces for persistence
│   └── service/            # Business logic layer
├── graph/
│   ├── generated/          # Auto-generated GraphQL code by gqlgen
│   ├── schema.graphqls     # GraphQL schema definitions
│   └── schema.resolvers.go # Resolver implementations
├── go.mod                  # Go module dependencies
├── gqlgen.yml              # gqlgen configuration
└── README.md
Key Principles:
internal packages are non-exportable outside the module, ensuring encapsulation.
GraphQL-specific logic is isolated in the graph package.
The cmd/server folder holds the executable, keeping application code separate from library code.
2. GraphQL Schema Design
The GraphQL API defines types, input objects, queries, and mutations.
Types
type Task {
  id: ID!
  title: String!
  status: String!
  projectID: ID!
}
Represents the Task domain object.
ID is the unique identifier.
status indicates task state (OPEN/CLOSED).
Input Types
input CreateTaskInput {
  title: String!
  projectID: ID!
}
Used as input for mutations.
Encourages future-proofing, allowing optional fields to be added without breaking clients.
Operations
Query: Fetch tasks (currently placeholder)
Mutation: Create a new task
type Mutation {
  createTask(input: CreateTaskInput!): Task!
}
Input objects encapsulate mutation parameters.
Returns a Task object after creation.
Concept: Using input objects instead of multiple arguments is a best practice for GraphQL mutation design.
3. Layered Architecture
The project implements three primary layers:
3.1 Resolver Layer
Located in graph/schema.resolvers.go.
Handles GraphQL requests.
Converts GraphQL input objects into service layer calls.
Avoids direct database access.
Example:
func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTaskInput) (*model.Task, error) {
    task := r.TaskService.CreateTask(input.Title, input.ProjectID)
    return task, nil
}
Concepts:
Resolvers are thin controllers: They orchestrate, not implement business logic.
Context (ctx) is passed for request-scoped data and cancellation.
3.2 Service Layer
Located in internal/service/task_service.go.
Encapsulates business logic.
Validates input and performs transformations.
Independent of GraphQL or database.
Example:
type TaskService struct {}

func (s *TaskService) CreateTask(title string, projectID string) *model.Task {
    return &model.Task{
        ID:        "1",
        Title:     title,
        Status:    "OPEN",
        ProjectID: projectID,
    }
}
Concepts:
Separation of concerns: Service layer does the "what" and "why", not the "how to store".
Future-proofing: Swapping memory storage for a database requires no changes in resolver logic.
3.3 Domain Models
Located in internal/model/task.go.
type Task struct {
    ID        string
    Title     string
    Status    string
    ProjectID string
}
Concepts:
Domain models represent real entities.
Used by service layer and GraphQL resolvers.
Enables consistency and type safety across layers.
4. Dependency Injection
Resolvers depend on services, not directly on repositories or DB.
Example:
srv := handler.NewDefaultServer(
    generated.NewExecutableSchema(
        generated.Config{
            Resolvers: &graph.Resolver{
                TaskService: &service.TaskService{},
            },
        },
    ),
)
Injecting TaskService ensures decoupled architecture.
Promotes unit testing and mocking.
5. GraphQL Execution Flow
Client sends a mutation query:
mutation {
  createTask(input: {title: "Learn Go", projectID: "101"}) {
    id
    title
    status
    projectID
  }
}
GraphQL server calls the resolver: CreateTask.
Resolver calls TaskService.CreateTask().
Service returns a Task object.
Resolver sends Task object back to the client.
Flow Diagram:
GraphQL Request
       │
       ▼
  Resolver Layer
       │
       ▼
  Service Layer
       │
       ▼
  (Future) Repository/Database
       │
       ▼
 Response to Client
