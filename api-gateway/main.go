package main

import (
	"api-gateway/config"
	"api-gateway/internal/client"
	"api-gateway/internal/handler"
	api "api-gateway/internal/server"
	"log"
	"os"
)

// @title           Food Delivery API Gateway
// @version         1.0
// @description     Gateway for microservices
// @host            localhost:8000
// @BasePath        /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := config.LoadConfig()

	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// gRPC clients
	authClient := client.InitAuthClient(cfg.AuthGRPCAddr)
	userClient := client.InitUserClient(cfg.UserGRPCAddr)
	orderClient := client.InitOrderClient(cfg.OrderGRPCAddr)
	restaurantClient := client.InitRestaurantClient(cfg.RestGRPCAddr)

	// Handlers
	authHandler := handler.NewAuthHandler(authClient)
	userHandler := handler.NewUserHandler(userClient)
	orderHandler := handler.NewOrderHandler(orderClient)
	restaurantHandler := handler.NewRestaurantHandler(restaurantClient)

	// HTTP Server
	server := api.NewServer(orderHandler, restaurantHandler, userHandler, authHandler)

	infoLog.Printf("starting API Gateway on port: %s", cfg.AppPort)
	server.Run(infoLog, errorLog)

}
