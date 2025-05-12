package handler

import (
	"google.golang.org/grpc"
	"user-service/internal/config"
	handlers "user-service/internal/handler/grpc"
	"user-service/internal/service/auth"
	"user-service/internal/service/user"
	authpb "user-service/proto/gen/authpb"
	userpb "user-service/proto/gen/userpb"
)

type Dependencies struct {
	Configs     config.Configs
	UserService *user.Service
	AuthService *auth.Service
}

type Configuration func(h *Handler) error
type Handler struct {
	dependencies Dependencies
	GRPC         *grpc.Server
}

func New(d Dependencies, configs ...Configuration) (h *Handler, err error) {
	// Create the handler
	h = &Handler{
		dependencies: d,
	}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the service into the configuration function
		if err = cfg(h); err != nil {
			return
		}
	}

	return
}

func WithGRPCHandler() Configuration {
	return func(h *Handler) error {
		h.GRPC = grpc.NewServer()

		// Регистрация gRPC-сервисов
		authpb.RegisterAuthServiceServer(h.GRPC, handlers.NewAuthHandler(h.dependencies.AuthService))
		userpb.RegisterUserServiceServer(h.GRPC, handlers.NewUserHandler(h.dependencies.UserService))

		return nil
	}
}
