package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	domain "user-service/internal/domain/user"
	"user-service/internal/service/user"
	"user-service/proto/gen/userpb"
)

type UserHandler struct {
	userpb.UnimplementedUserServiceServer
	service *user.Service
}

func NewUserHandler(s *user.Service) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUser(ctx context.Context, req *userpb.GetByIDRequest) (*userpb.UserResponse, error) {
	userEntity, err := h.service.GetByID(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "order not found: %v", err)
	}

	return &userpb.UserResponse{
		Id:    userEntity.ID,
		Name:  userEntity.Name,
		Email: userEntity.Email,
		Role:  userEntity.Role,
	}, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *userpb.UpdateRequest) (*userpb.Empty, error) {
	updateData := domain.Request{
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}

	err := h.service.Update(ctx, req.Id, updateData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update user: %v", err)
	}

	return &userpb.Empty{}, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *userpb.DeleteRequest) (*userpb.Empty, error) {
	err := h.service.Delete(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete order: %v", err)
	}

	return &userpb.Empty{}, nil
}
