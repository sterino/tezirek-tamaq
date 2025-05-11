package user

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, data Entity) (string, error)
	GetByEmail(ctx context.Context, email string) (Entity, error)
	GetByID(ctx context.Context, id string) (Entity, error)
	Update(ctx context.Context, id string, data Entity) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]Entity, error)
}
