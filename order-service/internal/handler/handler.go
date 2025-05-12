package handler

import (
	"google.golang.org/grpc"
	"order-service/internal/config"
	handlers "order-service/internal/handler/grpc"
	"order-service/internal/service/order"
)

type Dependencies struct {
	Configs      config.Configs
	OrderService *order.Service
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
		authpb.RegisterOrderServiceServer(h.GRPC, handlers.NewOrderHandler(h.dependencies.OrderService))

		return nil
	}
}
