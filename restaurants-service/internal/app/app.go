package app

import (
	"context"
	"flag"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"

	"restaurant-service/internal/config"
	"restaurant-service/internal/handler"
	"restaurant-service/internal/repository"
	"restaurant-service/internal/service/restaurant"
	"restaurant-service/pkg/log"
	"restaurant-service/pkg/server"
)

func Run() {
	logger := log.LoggerFromContext(context.Background())

	// Load configs
	cfg, err := config.New()
	if err != nil {
		logger.Fatal("failed to load config", zap.Error(err))
	}

	// Init repositories
	repositories, err := repository.New(
		repository.WithPostgresStore("postgres", cfg.POSTGRES.DSN))
	if err != nil {
		logger.Fatal("failed to init repositories", zap.Error(err))
	}
	defer repositories.Close()

	// Init services
	restaurantService, err := restaurant.New(
		restaurant.WithRestaurantRepository(repositories.Restaurant))
	if err != nil {
		logger.Fatal("failed to init restaurant service", zap.Error(err))
	}

	// Init handlers
	handlers, err := handler.New(
		handler.Dependencies{
			Configs:           *cfg,
			RestaurantService: restaurantService,
		},
		handler.WithGRPCHandler())
	if err != nil {
		logger.Fatal("failed to init handlers", zap.Error(err))
	}

	// Init gRPC server
	servers, err := server.New(
		server.WithGRPCServer(handlers.GRPC, cfg.APP.Port))
	if err != nil {
		logger.Fatal("failed to init server", zap.Error(err))
	}

	// Run server in goroutine
	go func() {
		if err := servers.Run(logger); err != nil {
			logger.Fatal("failed to run server", zap.Error(err))
		}
	}()
	logger.Info("gRPC server started", zap.String("port", cfg.APP.Port))

	// Graceful shutdown
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", 15*time.Second, "shutdown timeout")
	flag.Parse()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	if err := servers.Stop(ctx); err != nil {
		logger.Fatal("failed to stop server gracefully", zap.Error(err))
	}

	logger.Info("server shutdown complete")
}
