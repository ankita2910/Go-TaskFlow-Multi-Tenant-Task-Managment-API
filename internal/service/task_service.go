package service

import "go-graphql-taskflow/internal/graph/model"

type TaskService struct{}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (s *TaskService) CreateTask(title string, projectID string) *model.Task {
	return &model.Task{
		ID:    "1",
		Title: title,
		Status:    "OPEN",
        ProjectID: projectID,
	}
}