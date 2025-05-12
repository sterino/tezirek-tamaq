package app

import (
	"context"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	configs, err := config.New()
	if err != nil {
		logger.Error("ERR_INIT_CONFIGS", zap.Error(err))
		return
	}

	repositories, err := repository.New(
		repository.WithPostgresStore("postgres", "postgres://sterino:sterino@localhost:5432/postgres?sslmode=disable&search_path=public"))
	if err != nil {
		logger.Error("ERR_INIT_REPOSITORIES", zap.Error(err))
		return
	}
	defer repositories.Close()

	//caches, err := cache.New(
	//	cache.Dependencies{
	//		AuthorRepository: repositories.Author,
	//		BookRepository:   repositories.Book,
	//	},
	//	cache.WithMemoryStore())
	//if err != nil {
	//	logger.Error("ERR_INIT_CACHES", zap.Error(err))
	//	return
	//}
	//defer caches.Close()

	//paymentService, err := payment.New(
	//	payment.WithCurrencyClient(currencyClient))
	//if err != nil {
	//	logger.Error("ERR_INIT_PAYMENT_SERVICE", zap.Error(err))
	//	return
	//}

	authService, err := auth.New(
		auth.WithUserRepository(repositories.User))
	if err != nil {
		logger.Error("ERR_INIT_AUTH_SERVICE", zap.Error(err))
		return
	}

	userService, err := user.New(
		user.WithUserRepository(repositories.User))
	if err != nil {
		logger.Error("ERR_INIT_USER_SERVICE", zap.Error(err))
		return
	}

	handlers, err := handler.New(
		handler.Dependencies{
			Configs:     configs,
			AuthService: authService,
			UserService: userService,
		},
		handler.WithGRPCHandler())
	if err != nil {
		logger.Error("ERR_INIT_HANDLERS", zap.Error(err))
		return
	}

	servers, err := server.New(
		server.WithHTTPServer(handlers.GRPC, configs.APP.Port))
	if err != nil {
		logger.Error("ERR_INIT_SERVERS", zap.Error(err))
		return
	}

	// Run our server in a goroutine so that it doesn't block
	if err = servers.Run(logger); err != nil {
		logger.Error("ERR_RUN_SERVERS", zap.Error(err))
		return
	}
	logger.Info("http server started on http://localhost:" + configs.APP.Port) //"/swagger/index.html"

	// Graceful Shutdown
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the httpServer gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	quit := make(chan os.Signal, 1) // Create channel to signify a signal being sent

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel
	<-quit                                             // This blocks the main thread until an interrupt is received
	fmt.Println("gracefully shutting down...")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait until the timeout deadline
	if err = servers.Stop(ctx); err != nil {
		panic(err) // failure/timeout shutting down the httpServer gracefully
	}

	fmt.Println("running cleanup tasks...")
	// Your cleanup tasks go here

	fmt.Println("server was successful shutdown.")
}
