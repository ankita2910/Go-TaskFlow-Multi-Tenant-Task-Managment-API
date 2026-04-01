Go GraphQL Project: Conceptual Overview




1. Project Goal
Built a multi-tenant task management GraphQL API using Go.
The main objectives were:
Create and query tasks
Allow comments to be associated with tasks (planned)
Use services and repositories to separate business logic and data access
Use GraphQL resolvers to map API calls to services
Implement DataLoader middleware for efficient batching

2. Project Structure & Purpose of Each File

a) cmd/server/main.go – Entry Point / Server
Purpose: Bootstraps the entire application.
This file wires together all components—repositories, services, middleware, and GraphQL server.

b) internal/graph/schema.resolvers.go – GraphQL Resolvers
Purpose: Defines how GraphQL queries and mutations are handled.
Resolvers act as the bridge between GraphQL API and the business logic in services.

c) internal/graph/model – GraphQL Models
Purpose: Defines the data structures used in GraphQL API.
Models represent the shape of your GraphQL data.

d) internal/repository/comment_repository.go – Repository Layer
Purpose: Handles data access logic, separating storage from business logic.
Repositories isolate storage details, making your service layer independent of DB implementation.

e) internal/service/task_service.go – Service Layer
Purpose: Contains business logic.
Services prevent GraphQL resolvers from being cluttered with business rules.

f) internal/graph/loader – DataLoader Layer (Optional / Planned)
Purpose: Efficiently batch and cache database calls (like fetching comments for multiple tasks at once).
Avoids N+1 queries in GraphQL.

g) Middleware (DataLoaderMiddleware)
Wraps every HTTP request to inject loaders into context.
Ensures resolvers can access loaders for batched queries.

