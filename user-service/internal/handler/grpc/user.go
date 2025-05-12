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

func (h *UserHandler) CreateUser(ctx context.Context, req *userpb.UpdateRequest) (*userpb.UserResponse, error) {

	userEntity := domain.Request{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}

	createdUser, err := h.service.Create(ctx, userEntity)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	return &userpb.UserResponse{
		Id:    createdUser.ID,
		Name:  createdUser.Name,
		Email: createdUser.Email,
		Role:  createdUser.Role,
	}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *userpb.UserResponse) (*userpb.UserResponse, error) {
	userEntity, err := h.service.GetByID(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	return &userpb.UserResponse{
		Id:    userEntity.ID,
		Name:  userEntity.Name,
		Email: userEntity.Email,
		Role:  userEntity.Role,
	}, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *userpb.UpdateRequest) error {

	updateData := domain.Request{
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}

	err := h.service.Update(ctx, req.Id, updateData)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to update user: %v", err)
	}

	return nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *userpb.DeleteRequest) (*userpb.Empty, error) {
	err := h.service.Delete(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete user: %v", err)
	}

	return &userpb.Empty{}, nil
}
