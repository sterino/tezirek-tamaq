package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	"restaurant-service/internal/domain/restaurant"
	"restaurant-service/pkg/store"
)

type RestaurantRepository struct {
	db *sqlx.DB
}

func NewRestaurantRepository(db *sqlx.DB) *RestaurantRepository {
	return &RestaurantRepository{db: db}
}

func (r *RestaurantRepository) Create(ctx context.Context, data restaurant.Entity) (string, error) {
	query := `
		INSERT INTO restaurants (name, address, phone, order_ids, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	now := time.Now()
	args := []interface{}{
		data.Name,
		data.Address,
		data.Phone,
		pq.Array(data.OrderIDs),
		now,
		now,
	}

	var id string
	if err := r.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

func (r *RestaurantRepository) GetByID(ctx context.Context, id string) (restaurant.Entity, error) {
	query := `
		SELECT id, name, address, phone, order_ids, created_at, updated_at
		FROM restaurants
		WHERE id = $1`

	var entity restaurant.Entity
	if err := r.db.GetContext(ctx, &entity, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity, store.ErrorNotFound
		}
		return entity, err
	}
	return entity, nil
}

func (r *RestaurantRepository) Update(ctx context.Context, id string, data restaurant.Entity) error {
	sets, args := prepareUpdateArgs(data)
	if len(args) == 0 {
		return nil
	}

	args = append(args, id)
	sets = append(sets, "updated_at = NOW()")

	query := fmt.Sprintf("UPDATE restaurants SET %s WHERE id = $%d RETURNING id",
		strings.Join(sets, ", "), len(args))

	if err := r.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return store.ErrorNotFound
		}
		return err
	}
	return nil
}

func prepareUpdateArgs(data restaurant.Entity) ([]string, []interface{}) {
	var sets []string
	var args []interface{}

	if data.Name != "" {
		args = append(args, data.Name)
		sets = append(sets, fmt.Sprintf("name=$%d", len(args)))
	}
	if data.Address != "" {
		args = append(args, data.Address)
		sets = append(sets, fmt.Sprintf("address=$%d", len(args)))
	}
	if data.Phone != "" {
		args = append(args, data.Phone)
		sets = append(sets, fmt.Sprintf("phone=$%d", len(args)))
	}
	if len(data.OrderIDs) > 0 {
		args = append(args, pq.Array(data.OrderIDs))
		sets = append(sets, fmt.Sprintf("order_ids=$%d", len(args)))
	}

	return sets, args
}

func (r *RestaurantRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM restaurants WHERE id = $1 RETURNING id`

	if err := r.db.QueryRowContext(ctx, query, id).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return store.ErrorNotFound
		}
		return err
	}
	return nil
}

func (r *RestaurantRepository) List(ctx context.Context) ([]restaurant.Entity, error) {
	query := `
		SELECT id, name, address, phone, order_ids, created_at, updated_at
		FROM restaurants
		ORDER BY created_at DESC`

	var restaurants []restaurant.Entity
	if err := r.db.SelectContext(ctx, &restaurants, query); err != nil {
		return nil, err
	}
	return restaurants, nil
}
