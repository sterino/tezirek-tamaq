package order

type Request struct {
	UserID       string  `json:"user_id" validate:"required"`
	RestaurantID string  `json:"restaurant_id" validate:"required"`
	Items        []Item  `json:"items" validate:"required"`
	TotalPrice   float64 `json:"total_price" validate:"required"`
	Status       string  `json:"status"` // e.g., "pending", "paid", "delivered"
}

type Response struct {
	ID           string  `json:"id"`
	UserID       string  `json:"user_id"`
	RestaurantID string  `json:"restaurant_id"`
	Items        []Item  `json:"items"`
	TotalPrice   float64 `json:"total_price"`
	Status       string  `json:"status"`
}

func ParseFromEntity(e Entity) Response {
	return Response{
		ID:           e.ID,
		UserID:       e.UserID,
		RestaurantID: e.RestaurantID,
		Items:        e.Items,
		TotalPrice:   e.TotalPrice,
		Status:       e.Status,
	}
}

func ParseFromEntities(data []Entity) []Response {
	result := make([]Response, len(data))
	for i, entity := range data {
		result[i] = ParseFromEntity(entity)
	}
	return result
}
