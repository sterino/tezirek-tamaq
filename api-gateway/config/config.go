package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort       string
	JWTSecret     string
	AuthGRPCAddr  string
	UserGRPCAddr  string
	OrderGRPCAddr string
	RestGRPCAddr  string
	Timeout       time.Duration
}

func LoadConfig() Config {
	_ = godotenv.Load(".env") // Загружаем переменные окружения

	cfg := Config{
		AppPort:       getEnv("APP_PORT", "8000"),
		JWTSecret:     getEnv("JWT_SECRET", "super-secret"),
		AuthGRPCAddr:  getEnv("AUTH_GRPC_ADDR", "localhost:50051"),
		UserGRPCAddr:  getEnv("USER_GRPC_ADDR", "localhost:50051"),
		OrderGRPCAddr: getEnv("ORDER_GRPC_ADDR", "localhost:50052"),
		RestGRPCAddr:  getEnv("RESTAURANT_GRPC_ADDR", "localhost:50053"),
		Timeout:       10 * time.Second,
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	log.Printf("warning: %s not set, using default: %s", key, fallback)
	return fallback
}
