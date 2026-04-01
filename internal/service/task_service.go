package service

import (
    "context"

    "go-graphql-taskflow/internal/graph/model"
    "go-graphql-taskflow/internal/repository"

    "github.com/google/uuid"
)

type TaskService struct {
    Repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
    return &TaskService{
        Repo: repo,
    }
}

func (s *TaskService) CreateTask(title, projectID string) (*model.Task, error) {
    task := &model.Task{
        ID:        uuid.NewString(),
        Title:     title,
        Status:    "OPEN",
        ProjectID: projectID,
    }
    err := s.Repo.CreateTask(context.Background(), task)
    return task, err
}