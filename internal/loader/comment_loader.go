package loader

import (
	"context"
	"time"

	"github.com/graph-gophers/dataloader/v7"
	"go-graphql-taskflow/internal/graph/model"
	"go-graphql-taskflow/internal/service"
)

type CommentLoader struct {
	Loader *dataloader.Loader[string, []*model.Comment]
}

func NewCommentLoader(commentService *service.CommentService) *CommentLoader {

	batchFn := func(ctx context.Context, taskIDs []string) []*dataloader.Result[[]*model.Comment] {

		// SINGLE DB CALL
		commentMap, err := commentService.GetCommentsByTaskIDs(ctx, taskIDs)

		results := make([]*dataloader.Result[[]*model.Comment], len(taskIDs))

		for i, id := range taskIDs {
			results[i] = &dataloader.Result[[]*model.Comment]{
				Data:  commentMap[id],
				Error: err,
			}
		}

		return results
	}

	loader := dataloader.NewBatchedLoader(
		batchFn,
		dataloader.WithWait[string, []*model.Comment](2*time.Millisecond),
	)

	return &CommentLoader{
		Loader: loader,
	}
}