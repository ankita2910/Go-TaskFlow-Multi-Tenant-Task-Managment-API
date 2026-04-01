package service

import (
	"context"

	"go-graphql-taskflow/internal/graph/model"
	"go-graphql-taskflow/internal/repository"
)

type CommentService struct {
	Repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) *CommentService {
	return &CommentService{
		Repo: repo,
	}
}

func (s *CommentService) GetCommentsByTaskIDs(
	ctx context.Context,
	taskIDs []string,
) (map[string][]*model.Comment, error) {

	return s.Repo.BatchGetComments(ctx, taskIDs)
}