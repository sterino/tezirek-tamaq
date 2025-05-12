package order

import "time"

type Entity struct {
	ID           string    `db:"id"`
	UserID       string    `db:"user_id"`
	RestaurantID string    `db:"restaurant_id"`
	Items        []Item    `db:"items"` // stored as JSONB
	TotalPrice   float64   `db:"total_price"`
	Status       string    `db:"status"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type Item struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func NewCreate(req Request) Entity {
	return Entity{
		UserID:       req.UserID,
		RestaurantID: req.RestaurantID,
		Items:        req.Items,
		TotalPrice:   req.TotalPrice,
		Status:       req.Status,
	}
}

func NewUpdate(req Request) Entity {
	return Entity{
		UserID:       req.UserID,
		RestaurantID: req.RestaurantID,
		Items:        req.Items,
		TotalPrice:   req.TotalPrice,
		Status:       req.Status,
	}
}
