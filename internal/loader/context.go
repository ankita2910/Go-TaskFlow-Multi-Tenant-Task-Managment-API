package loader

import "context"

type ctxKey string

const loadersKey ctxKey = "dataloaders"

type Loaders struct {
	CommentLoader *CommentLoader
}

func InjectLoaders(ctx context.Context, loaders *Loaders) context.Context {
	return context.WithValue(ctx, loadersKey, loaders)
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}