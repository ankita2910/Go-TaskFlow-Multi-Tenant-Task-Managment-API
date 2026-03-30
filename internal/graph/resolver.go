package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.
import (
    "go-graphql-taskflow/internal/service"
)
type Resolver struct{
	TaskService *service.TaskService
}
