package app

import (
	"context"
	"flag"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"

	"order-service/internal/config"
	"order-service/internal/handler"
	"order-service/internal/repository"
	"order-service/internal/service/order"
	"order-service/pkg/log"
	"order-service/pkg/server"
)

func Run() {
	logger := log.LoggerFromContext(context.Background())

	// Load config
	cfg, err := config.New()
	if err != nil {
		logger.Fatal("failed to load config", zap.Error(err))
	}

	// Init repository
	repositories, err := repository.New(
		repository.WithPostgresStore("postgres", cfg.POSTGRES.DSN))
	if err != nil {
		logger.Fatal("failed to init repositories", zap.Error(err))
	}
	defer repositories.Close()

	// Init service
	orderService, err := order.New(
		order.WithOrderRepository(repositories.Order))
	if err != nil {
		logger.Fatal("failed to init order service", zap.Error(err))
	}

	// Init handler
	handlers, err := handler.New(
		handler.Dependencies{
			Configs:      *cfg,
			OrderService: orderService,
		},
		handler.WithGRPCHandler())
	if err != nil {
		logger.Fatal("failed to init handlers", zap.Error(err))
	}

	// Init server
	servers, err := server.New(
		server.WithGRPCServer(handlers.GRPC, cfg.APP.Port))
	if err != nil {
		logger.Fatal("failed to init server", zap.Error(err))
	}

	// Run server
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
