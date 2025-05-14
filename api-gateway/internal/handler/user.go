package handler

import (
	"api-gateway/pkg/response"
	"api-gateway/proto/gen/userpb"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	client userpb.UserServiceClient
}

func NewUserHandler(client userpb.UserServiceClient) *UserHandler {
	return &UserHandler{client: client}
}

// GetUserByID godoc
// @Summary     Get user by ID
// @Description Get a single user by their ID
// @Tags        users
// @Security    BearerAuth
// @Produce     json
// @Param       id path string true "User ID"
// @Success     200 {object} response.UserResponse
// @Failure     404 {object} response.ErrorResponse
// @Router      /user/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.client.GetByID(context.Background(), &userpb.GetByIDRequest{Id: id})
	if err != nil {
		response.NewError(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.FromUserProto(res))
}

// ListUsers godoc
// @Summary     List all users
// @Description Retrieve all registered users
// @Tags        users
// @Security    BearerAuth
// @Produce     json
// @Success     200 {array} response.UserResponse
// @Failure     500 {object} response.ErrorResponse
// @Router      /user/all [get]
func (h *UserHandler) List(c *gin.Context) {
	res, err := h.client.List(context.Background(), &userpb.Empty{})
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, "failed to list users")
		return
	}

	users := make([]response.UserResponse, 0, len(res.Users))
	for _, u := range res.Users {
		users = append(users, response.FromUserProto(u))
	}

	c.JSON(http.StatusOK, users)
}
