package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	domain "user-service/internal/domain/auth"
	"user-service/internal/domain/user"
	"user-service/internal/service/auth"
	"user-service/proto/gen/authpb"
)

type AuthHandler struct {
	authpb.UnimplementedAuthServiceServer
	service *auth.Service
}

func NewAuthHandler(s *auth.Service) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.TokenResponse, error) {
	token, _, err := h.service.Login(ctx, domain.Request{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %v", err)
	}
	return &authpb.TokenResponse{Token: token}, nil
}

func (h *AuthHandler) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.UserIDResponse, error) {
	// Реализуйте логику регистрации пользователя здесь
	// Например:
	userID, err := h.service.Register(ctx, user.Request{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		Role:     "order",
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "registration failed: %v", err)
	}

	return &authpb.UserIDResponse{Id: userID}, nil
}
