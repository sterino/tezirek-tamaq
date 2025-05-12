package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"order-service/internal/domain/order"
	"order-service/pkg/store"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(ctx context.Context, data order.Entity) (string, error) {
	query := `
		INSERT INTO orders (user_id, restaurant_id, items, total_price, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`

	itemsJSON, err := json.Marshal(data.Items)
	if err != nil {
		return "", err
	}

	now := time.Now()
	args := []interface{}{data.UserID, data.RestaurantID, itemsJSON, data.TotalPrice, data.Status, now, now}

	var id string
	if err := r.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

func (r *OrderRepository) GetByID(ctx context.Context, id string) (order.Entity, error) {
	query := `SELECT id, user_id, restaurant_id, items, total_price, status, created_at, updated_at FROM orders WHERE id = $1`

	var raw struct {
		order.Entity
		ItemsRaw []byte `db:"items"`
	}

	if err := r.db.GetContext(ctx, &raw, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return order.Entity{}, store.ErrorNotFound
		}
		return order.Entity{}, err
	}

	if err := json.Unmarshal(raw.ItemsRaw, &raw.Items); err != nil {
		return order.Entity{}, err
	}

	return raw.Entity, nil
}

func (r *OrderRepository) ListByUserID(ctx context.Context, userID string) ([]order.Entity, error) {
	query := `SELECT id, user_id, restaurant_id, items, total_price, status, created_at, updated_at FROM orders WHERE user_id = $1`

	return r.selectAndDecode(ctx, query, userID)
}

func (r *OrderRepository) ListByRestaurantID(ctx context.Context, restaurantID string) ([]order.Entity, error) {
	query := `SELECT id, user_id, restaurant_id, items, total_price, status, created_at, updated_at FROM orders WHERE restaurant_id = $1`

	return r.selectAndDecode(ctx, query, restaurantID)
}

func (r *OrderRepository) List(ctx context.Context) ([]order.Entity, error) {
	query := `SELECT id, user_id, restaurant_id, items, total_price, status, created_at, updated_at FROM orders ORDER BY created_at DESC`

	return r.selectAndDecode(ctx, query)
}

func (r *OrderRepository) Update(ctx context.Context, id string, data order.Entity) error {
	sets, args := prepareUpdateArgs(data)
	if len(args) == 0 {
		return nil
	}

	args = append(args, id)
	sets = append(sets, "updated_at = NOW()")

	query := fmt.Sprintf("UPDATE orders SET %s WHERE id = $%d RETURNING id", strings.Join(sets, ", "), len(args))

	if err := r.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return store.ErrorNotFound
		}
		return err
	}
	return nil
}

func (r *OrderRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM orders WHERE id = $1 RETURNING id`

	if err := r.db.QueryRowContext(ctx, query, id).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return store.ErrorNotFound
		}
		return err
	}
	return nil
}

func (r *OrderRepository) selectAndDecode(ctx context.Context, query string, args ...interface{}) ([]order.Entity, error) {
	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []order.Entity
	for rows.Next() {
		var raw struct {
			order.Entity
			ItemsRaw []byte `db:"items"`
		}
		if err := rows.StructScan(&raw); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(raw.ItemsRaw, &raw.Items); err != nil {
			return nil, err
		}
		results = append(results, raw.Entity)
	}
	return results, nil
}

func prepareUpdateArgs(data order.Entity) ([]string, []interface{}) {
	var sets []string
	var args []interface{}

	if data.Status != "" {
		args = append(args, data.Status)
		sets = append(sets, fmt.Sprintf("status=$%d", len(args)))
	}
	if len(data.Items) > 0 {
		jsonData, _ := json.Marshal(data.Items)
		args = append(args, jsonData)
		sets = append(sets, fmt.Sprintf("items=$%d", len(args)))
	}
	if data.TotalPrice > 0 {
		args = append(args, data.TotalPrice)
		sets = append(sets, fmt.Sprintf("total_price=$%d", len(args)))
	}
	return sets, args
}
