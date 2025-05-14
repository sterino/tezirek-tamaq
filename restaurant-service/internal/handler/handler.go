package handler

import (
	"google.golang.org/grpc"
	"restaurant-service/internal/config"
	handlers "restaurant-service/internal/handler/grpc"
	"restaurant-service/internal/service/restaurant"
	"restaurant-service/proto/gen/restaurantpb"
)

type Dependencies struct {
	Configs           config.Configs
	RestaurantService *restaurant.Service
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
		restaurantpb.RegisterRestaurantServiceServer(h.GRPC, handlers.NewRestaurantHandler(h.dependencies.RestaurantService))

		return nil
	}
}
