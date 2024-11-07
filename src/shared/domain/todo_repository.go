package domain

import (
	"context"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *Todo) error
	Update(ctx context.Context, todo *Todo) (*Todo, error)
	Delete(ctx context.Context, todoId string) error
	Get(ctx context.Context, todoId string) (*Todo, error)
}
