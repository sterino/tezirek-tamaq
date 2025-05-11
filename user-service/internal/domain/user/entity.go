package user

import "time"

type Entity struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"` // не экспортируем наружу
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func New(req Request, hashPassword string) Entity {
	return Entity{
		Name:     req.Name,
		Email:    req.Email,
		Role:     req.Role,
		Password: hashPassword,
	}
}
