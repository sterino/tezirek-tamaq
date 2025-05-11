package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"

	"user-service/internal/domain/user"
	"user-service/pkg/store"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, data user.Entity) (string, error) {
	query := `
		INSERT INTO users (name, email, password, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	now := time.Now()
	args := []interface{}{data.Name, data.Email, data.Password, data.Role, now, now}

	var id string
	if err := r.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (user.Entity, error) {
	query := `SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE email=$1`

	var entity user.Entity
	if err := r.db.GetContext(ctx, &entity, query, email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity, store.ErrorNotFound
		}
		return entity, err
	}
	return entity, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (user.Entity, error) {
	query := `SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE id=$1`

	var entity user.Entity
	if err := r.db.GetContext(ctx, &entity, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity, store.ErrorNotFound
		}
		return entity, err
	}
	return entity, nil
}

func (r *UserRepository) Update(ctx context.Context, id string, data user.Entity) error {
	sets, args := r.prepareUpdateArgs(data)
	if len(args) == 0 {
		return nil
	}

	args = append(args, id)
	sets = append(sets, fmt.Sprintf("updated_at=NOW()"))
	query := fmt.Sprintf("UPDATE users SET %s WHERE id=$%d RETURNING id", strings.Join(sets, ", "), len(args))

	if err := r.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return store.ErrorNotFound
		}
		return err
	}
	return nil
}

func (r *UserRepository) prepareUpdateArgs(data user.Entity) ([]string, []interface{}) {
	var sets []string
	var args []interface{}

	if data.Name != "" {
		args = append(args, data.Name)
		sets = append(sets, fmt.Sprintf("name=$%d", len(args)))
	}
	if data.Email != "" {
		args = append(args, data.Email)
		sets = append(sets, fmt.Sprintf("email=$%d", len(args)))
	}
	if data.Password != "" {
		args = append(args, data.Password)
		sets = append(sets, fmt.Sprintf("password=$%d", len(args)))
	}
	if data.Role != "" {
		args = append(args, data.Role)
		sets = append(sets, fmt.Sprintf("role=$%d", len(args)))
	}

	return sets, args
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id=$1 RETURNING id`

	if err := r.db.QueryRowContext(ctx, query, id).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return store.ErrorNotFound
		}
		return err
	}
	return nil
}

func (r *UserRepository) List(ctx context.Context) ([]user.Entity, error) {
	query := `SELECT id, name, email, password, role, created_at, updated_at FROM users ORDER BY created_at DESC`

	var users []user.Entity
	if err := r.db.SelectContext(ctx, &users, query); err != nil {
		return nil, err
	}
	return users, nil
}
