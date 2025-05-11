package handler

import "tezirek_tamaq/order-service/internal/config"

type Dependencies struct {
	Configs config.Configs
}

// tyoe Configuration func(h *Handler) error
type Handler struct {
	dependencies Dependencies
	//HTTP         *gin.Engine
}
