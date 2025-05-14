package app

import (
	"context"
	"flag"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user-service/pkg/jwt"

	"user-service/internal/config"
	"user-service/internal/handler"
	"user-service/internal/repository"
	"user-service/internal/service/auth"
	"user-service/internal/service/user"
	"user-service/pkg/log"
	"user-service/pkg/server"
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
		repository.WithPostgresStore("postgres", cfg.Postgres.DSN))
	if err != nil {
		logger.Fatal("failed to init repositories", zap.Error(err))
	}
	defer repositories.Close()

	// Init services
	secretKey := jwt.ProvideSecretKey()
	authService, err := auth.New(
		auth.WithUserRepository(repositories.User, secretKey))
	if err != nil {
		logger.Fatal("failed to init auth service", zap.Error(err))
	}

	userService, err := user.New(
		user.WithUserRepository(repositories.User))
	if err != nil {
		logger.Fatal("failed to init user service", zap.Error(err))
	}

	// Init handlers
	handlers, err := handler.New(
		handler.Dependencies{
			Configs:     *cfg,
			AuthService: authService,
			UserService: userService,
		},
		handler.WithGRPCHandler())
	if err != nil {
		logger.Fatal("failed to init handlers", zap.Error(err))
	}

	// Init server
	servers, err := server.New(
		server.WithGRPCServer(handlers.GRPC, cfg.App.Port))
	if err != nil {
		logger.Fatal("failed to init server", zap.Error(err))
	}

	// Run server in goroutine
	go func() {
		if err := servers.Run(logger); err != nil {
			logger.Fatal("failed to run server", zap.Error(err))
		}
	}()
	logger.Info("gRPC server started", zap.String("port", cfg.App.Port))

	// Graceful shutdown setup
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
		logger.Fatal("failed to gracefully stop server", zap.Error(err))
	}

	logger.Info("server shutdown complete")
}
