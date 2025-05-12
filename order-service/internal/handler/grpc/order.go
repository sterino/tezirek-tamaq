package grpc

import (
	"context"
	domain "order-service/internal/domain/order"
	"order-service/internal/service/order"
	"order-service/proto/gen/orderpb"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrderHandler struct {
	orderpb.UnimplementedOrderServiceServer
	service *order.Service
}

func NewOrderHandler(s *order.Service) *OrderHandler {
	return &OrderHandler{service: s}
}

func (h *OrderHandler) Create(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.OrderResponse, error) {
	res, err := h.service.Create(ctx, domain.Request{
		UserID:       req.GetUserId(),
		RestaurantID: req.GetRestaurantId(),
		Items:        toDomainItems(req.GetItems()),
		TotalPrice:   req.GetTotalPrice(),
		Status:       req.GetStatus(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create order: %v", err)
	}
	return toProtoResponse(res), nil
}

func (h *OrderHandler) GetByID(ctx context.Context, req *orderpb.GetOrderByIDRequest) (*orderpb.OrderResponse, error) {
	res, err := h.service.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "order not found: %v", err)
	}
	return toProtoResponse(res), nil
}

func (h *OrderHandler) Update(ctx context.Context, req *orderpb.UpdateOrderRequest) (*emptypb.Empty, error) {
	err := h.service.Update(ctx, req.GetId(), domain.Request{
		Items:      toDomainItems(req.GetItems()),
		TotalPrice: req.GetTotalPrice(),
		Status:     req.GetStatus(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update order: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func (h *OrderHandler) Delete(ctx context.Context, req *orderpb.DeleteOrderRequest) (*emptypb.Empty, error) {
	if err := h.service.Delete(ctx, req.GetId()); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete order: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func (h *OrderHandler) List(ctx context.Context, _ *emptypb.Empty) (*orderpb.ListOrdersResponse, error) {
	list, err := h.service.List(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list orders: %v", err)
	}

	var protoList []*orderpb.OrderResponse
	for _, o := range list {
		protoList = append(protoList, toProtoResponse(o))
	}
	return &orderpb.ListOrdersResponse{Orders: protoList}, nil
}

func (h *OrderHandler) ListByUser(ctx context.Context, req *orderpb.ListByUserRequest) (*orderpb.ListOrdersResponse, error) {
	list, err := h.service.ListByUser(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list orders by user: %v", err)
	}

	var protoList []*orderpb.OrderResponse
	for _, o := range list {
		protoList = append(protoList, toProtoResponse(o))
	}
	return &orderpb.ListOrdersResponse{Orders: protoList}, nil
}

func (h *OrderHandler) ListByRestaurant(ctx context.Context, req *orderpb.ListByRestaurantRequest) (*orderpb.ListOrdersResponse, error) {
	list, err := h.service.ListByRestaurant(ctx, req.GetRestaurantId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list orders by restaurant: %v", err)
	}

	var protoList []*orderpb.OrderResponse
	for _, o := range list {
		protoList = append(protoList, toProtoResponse(o))
	}
	return &orderpb.ListOrdersResponse{Orders: protoList}, nil
}

// ==== HELPERS ====

func toProtoResponse(o domain.Response) *orderpb.OrderResponse {
	return &orderpb.OrderResponse{
		Id:           o.ID,
		UserId:       o.UserID,
		RestaurantId: o.RestaurantID,
		Items:        toProtoItems(o.Items),
		TotalPrice:   o.TotalPrice,
		Status:       o.Status,
	}
}

func toProtoItems(items []domain.Item) []*orderpb.Item {
	protoItems := make([]*orderpb.Item, len(items))
	for i, item := range items {
		protoItems[i] = &orderpb.Item{
			Name:     item.Name,
			Quantity: int32(item.Quantity),
			Price:    item.Price,
		}
	}
	return protoItems
}

func toDomainItems(protoItems []*orderpb.Item) []domain.Item {
	items := make([]domain.Item, len(protoItems))
	for i, item := range protoItems {
		items[i] = domain.Item{
			Name:     item.GetName(),
			Quantity: int(item.GetQuantity()),
			Price:    item.GetPrice(),
		}
	}
	return items
}

func toTimestamp(t time.Time) *timestamppb.Timestamp {
	if t.IsZero() {
		return nil
	}
	return timestamppb.New(t)
}
