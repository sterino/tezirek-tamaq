package user

import "time"

type Entity struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewUpdate(req Request) Entity {
	return Entity{
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}
}

func NewCreate(req Request, hashedPassword string) Entity {
	return Entity{
		Name:     req.Name,
		Email:    req.Email,
		Role:     req.Role,
		Password: hashedPassword,
	}
}
