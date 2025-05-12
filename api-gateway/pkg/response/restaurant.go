package response

import (
	"api-gateway/proto/gen/restaurantpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type RestaurantResponse struct {
	ID        string   `json:"id" example:"resto-456"`
	Name      string   `json:"name" example:"Best Pizza"`
	Phone     string   `json:"description"`
	Location  string   `json:"location" example:"123 Main St"`
	Orders    []string `json:"orders" example:"[\"order-1\", \"order-2\"]"`
	CreatedAt string   `json:"created_at" example:"2024-05-12T12:00:00Z"`
	UpdatedAt string   `json:"updated_at" example:"2024-05-12T13:00:00Z"`
}

func FromRestaurantProto(pb *restaurantpb.RestaurantResponse) RestaurantResponse {
	return RestaurantResponse{
		ID:        pb.GetId(),
		Name:      pb.GetName(),
		Phone:     pb.GetPhone(),
		Location:  pb.GetAddress(),
		Orders:    pb.GetOrderIds(),
		CreatedAt: formatTimestamp(pb.GetCreatedAt()),
		UpdatedAt: formatTimestamp(pb.GetUpdatedAt()),
	}
}

func formatTimestamp(ts *timestamppb.Timestamp) string {
	if ts == nil {
		return ""
	}
	return ts.AsTime().Format(time.RFC3339)
}
