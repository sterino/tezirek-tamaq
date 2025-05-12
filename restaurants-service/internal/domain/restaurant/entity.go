package restaurant

import "time"

type Entity struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Address   string    `db:"address"`
	Phone     string    `db:"phone"`
	OrderIDs  []string  `db:"order_ids"` // Предполагается, что используется тип, поддерживающий массив строк, например, PostgreSQL
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewCreate(req Request) Entity {
	return Entity{
		Name:     req.Name,
		Address:  req.Address,
		Phone:    req.Phone,
		OrderIDs: req.OrderIDs,
	}
}

func NewUpdate(req Request) Entity {
	return Entity{
		Name:     req.Name,
		Address:  req.Address,
		Phone:    req.Phone,
		OrderIDs: req.OrderIDs,
	}
}
