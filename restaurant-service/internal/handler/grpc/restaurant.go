package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	domain "restaurant-service/internal/domain/restaurant"
	"restaurant-service/internal/service/restaurant"
	"restaurant-service/proto/gen/restaurantpb"
	"time"
)

type RestaurantHandler struct {
	restaurantpb.UnimplementedRestaurantServiceServer
	service *restaurant.Service
}

func NewRestaurantHandler(s *restaurant.Service) *RestaurantHandler {
	return &RestaurantHandler{service: s}
}

func (h *RestaurantHandler) Create(ctx context.Context, req *restaurantpb.CreateRestaurantRequest) (*restaurantpb.RestaurantResponse, error) {
	res, err := h.service.Create(ctx, domain.Request{
		Name:     req.GetName(),
		Address:  req.GetAddress(),
		Phone:    req.GetPhone(),
		OrderIDs: req.GetOrderIds(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create restaurant: %v", err)
	}
	return toProtoResponse(res), nil
}

func (h *RestaurantHandler) GetByID(ctx context.Context, req *restaurantpb.GetRestaurantByIDRequest) (*restaurantpb.RestaurantResponse, error) {
	res, err := h.service.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "restaurant not found: %v", err)
	}
	return toProtoResponse(res), nil
}

func (h *RestaurantHandler) Update(ctx context.Context, req *restaurantpb.UpdateRestaurantRequest) (*emptypb.Empty, error) {
	err := h.service.Update(ctx, req.GetId(), domain.Request{
		Name:     req.GetName(),
		Address:  req.GetAddress(),
		Phone:    req.GetPhone(),
		OrderIDs: req.GetOrderIds(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update restaurant: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func (h *RestaurantHandler) Delete(ctx context.Context, req *restaurantpb.DeleteRestaurantRequest) (*emptypb.Empty, error) {
	if err := h.service.Delete(ctx, req.GetId()); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete restaurant: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func (h *RestaurantHandler) List(ctx context.Context, _ *emptypb.Empty) (*restaurantpb.ListRestaurantsResponse, error) {
	list, err := h.service.List(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list restaurants: %v", err)
	}

	var protoList []*restaurantpb.RestaurantResponse
	for _, r := range list {
		protoList = append(protoList, toProtoResponse(r))
	}
	return &restaurantpb.ListRestaurantsResponse{Restaurants: protoList}, nil
}

// Helper: convert internal Response → gRPC
func toProtoResponse(r domain.Response) *restaurantpb.RestaurantResponse {
	return &restaurantpb.RestaurantResponse{
		Id:       r.ID,
		Name:     r.Name,
		Address:  r.Address,
		Phone:    r.Phone,
		OrderIds: r.OrderIDs,
	}
}

func toTimestamp(t time.Time) *timestamppb.Timestamp {
	if t.IsZero() {
		return nil
	}
	return timestamppb.New(t)
}
