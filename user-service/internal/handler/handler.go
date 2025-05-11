package handler

import (
	"user-service/internal/config"
	"user-service/internal/service/auth"
	"user-service/internal/service/user"
)

type Dependencies struct {
	Configs config.Configs
}

type Configuration func(h *Handler) error
type Handler struct {
	dependencies Dependencies
	userService  *user.Service
	authService  *auth.Service
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
