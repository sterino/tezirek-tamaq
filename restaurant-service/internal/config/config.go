package config

import (
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	defaultAppMode    = "dev"
	defaultAppPort    = "50053" // можно поменять для restaurant
	defaultAppPath    = "/"
	defaultAppTimeout = 60 * time.Second

	defaultTokenSalt    = "IP03O5Ekg91g5jw=="
	defaultTokenExpires = 3600 * time.Second
)

type (
	Configs struct {
		APP      AppConfig
		TOKEN    TokenConfig
		POSTGRES StoreConfig
	}

	AppConfig struct {
		Mode    string        `envconfig:"MODE" default:"dev"`
		Port    string        `envconfig:"PORT" default:"50053"`
		Path    string        `envconfig:"PATH" default:"/"`
		Timeout time.Duration `envconfig:"TIMEOUT" default:"60s"`
	}

	TokenConfig struct {
		Salt    string        `envconfig:"SALT" default:"IP03O5Ekg91g5jw=="`
		Expires time.Duration `envconfig:"EXPIRES" default:"3600s"`
	}

	StoreConfig struct {
		DSN string `envconfig:"DSN" required:"true"`
	}
)

// New загружает переменные окружения и возвращает сконфигурированную структуру
func New() (*Configs, error) {
	_ = godotenv.Load(filepath.Join(".", ".env"))

	cfg := &Configs{}

	if err := envconfig.Process("APP", &cfg.APP); err != nil {
		return nil, err
	}
	if err := envconfig.Process("TOKEN", &cfg.TOKEN); err != nil {
		return nil, err
	}
	if err := envconfig.Process("POSTGRES", &cfg.POSTGRES); err != nil {
		return nil, err
	}

	return cfg, nil
}
