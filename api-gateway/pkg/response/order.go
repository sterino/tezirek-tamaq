package response

import (
	"api-gateway/proto/gen/orderpb"
	"time"
)

type OrderResponse struct {
	Id           string  `json:"id" example:"123"`
	UserId       string  `json:"user_id"`
	RestaurantId string  `json:"restaurant_id"`
	Status       string  `json:"status"`
	Price        float64 `json:"price"`
	CreatedAt    string  `json:"created_at" example:"2024-05-12T12:00:00Z"`
	UpdatedAt    string  `json:"updated_at" example:"2024-05-12T13:00:00Z"`
}

func FromOrderProto(o *orderpb.OrderResponse) OrderResponse {
	var createdAt, updatedAt string
	if o.GetCreatedAt() != nil {
		createdAt = o.GetCreatedAt().AsTime().Format(time.RFC3339)
	}
	if o.GetUpdatedAt() != nil {
		updatedAt = o.GetUpdatedAt().AsTime().Format(time.RFC3339)
	}

	return OrderResponse{
		Id:           o.GetId(),
		UserId:       o.GetUserId(),
		RestaurantId: o.GetRestaurantId(),
		Status:       o.GetStatus(),
		Price:        o.GetTotalPrice(),
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}
