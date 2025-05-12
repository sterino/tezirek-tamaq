package handler

import (
	"api-gateway/pkg/response"
	"api-gateway/proto/gen/authpb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	client authpb.AuthServiceClient
}

func NewAuthHandler(client authpb.AuthServiceClient) *AuthHandler {
	return &AuthHandler{client: client}
}

// Login godoc
// @Summary     Авторизация
// @Description Аутентификация пользователя и получение токена
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       request body authpb.LoginRequest true "Данные для входа"
// @Success     200 {object} response.AuthResponse
// @Failure     400 {object} response.ErrorResponse
// @Failure     401 {object} response.ErrorResponse
// @Router      /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req authpb.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewError(c, http.StatusBadRequest, "invalid login data")
		return
	}

	res, err := h.client.Login(c.Request.Context(), &req)
	if err != nil {
		response.NewError(c, http.StatusUnauthorized, "invalid credentials")
		return
	}

	c.JSON(http.StatusOK, response.FromAuthProto(res))
}

// Register godoc
// @Summary     Регистрация
// @Description Создание нового пользователя
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       request body authpb.RegisterRequest true "Данные для регистрации"
// @Success     201 {object} authpb.UserIDResponse
// @Failure     400 {object} response.ErrorResponse
// @Failure     500 {object} response.ErrorResponse
// @Router      /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req authpb.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewError(c, http.StatusBadRequest, "invalid registration data")
		return
	}

	res, err := h.client.Register(c.Request.Context(), &req)
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, "registration failed")
		return
	}

	c.JSON(http.StatusCreated, res)
}
