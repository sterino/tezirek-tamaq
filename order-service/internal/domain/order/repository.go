package order

import "context"

type Repository interface {
	Create(ctx context.Context, data Entity) (string, error)
	GetByID(ctx context.Context, id string) (Entity, error)
	ListByUserID(ctx context.Context, userID string) ([]Entity, error)
	ListByRestaurantID(ctx context.Context, restaurantID string) ([]Entity, error)
	Update(ctx context.Context, id string, data Entity) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]Entity, error)
}
