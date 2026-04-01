package repository

import (
	"context"

	"go-graphql-taskflow/internal/graph/model"
)

type CommentRepository interface {
	BatchGetComments(
		ctx context.Context,
		taskIDs []string,
	) (map[string][]*model.Comment, error)
}
type TaskRepository interface {
	CreateTask(ctx context.Context, task *model.Task) error
	GetTasks(ctx context.Context, projectID string) ([]*model.Task, error)
}

type taskRepoImpl struct {
	tasks map[string]*model.Task
}
func (r *taskRepoImpl) CreateTask(ctx context.Context, task *model.Task) error {
	if r.tasks == nil {
		r.tasks = make(map[string]*model.Task)
	}
	r.tasks[task.ID] = task
	return nil
}

func (r *taskRepoImpl) GetTasks(ctx context.Context, projectID string) ([]*model.Task, error) {
	result := []*model.Task{}
	for _, t := range r.tasks {
		if t.ProjectID == projectID {
			result = append(result, t)
		}
	}
	return result, nil
}

// --- constructor function ---
func NewTaskRepository() TaskRepository {
	return &taskRepoImpl{
		tasks: make(map[string]*model.Task),
	}
}

type commentRepoImpl struct{}

func (r *commentRepoImpl) BatchGetComments(ctx context.Context, taskIDs []string) (map[string][]*model.Comment, error) {
	return make(map[string][]*model.Comment), nil
}

// Constructor to match what main.go expects
func NewCommentRepository() CommentRepository {
	return &commentRepoImpl{}
}