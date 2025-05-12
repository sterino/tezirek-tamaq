package handler

import (
	"api-gateway/pkg/response"
	"api-gateway/proto/gen/restaurantpb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RestaurantHandler struct {
	client restaurantpb.RestaurantServiceClient
}

func NewRestaurantHandler(client restaurantpb.RestaurantServiceClient) *RestaurantHandler {
	return &RestaurantHandler{client: client}
}

// CreateRestaurant godoc
// @Summary     Create a new restaurant
// @Description Create and return a restaurant
// @Tags        restaurants
// @Security BearerAuth
// @Accept      json
// @Produce     json
// @Param       request body restaurantpb.CreateRestaurantRequest true "Restaurant info"
// @Success     201 {object} response.RestaurantResponse
// @Failure     400 {object} response.ErrorResponse
// @Failure     500 {object} response.ErrorResponse
// @Router      /restaurant/create [post]
func (h *RestaurantHandler) Create(c *gin.Context) {
	var req restaurantpb.CreateRestaurantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	res, err := h.client.Create(c.Request.Context(), &req)
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, "failed to create restaurant")
		return
	}

	c.JSON(http.StatusCreated, response.FromRestaurantProto(res))
}

// GetRestaurantByID godoc
// @Summary     Get restaurant by ID
// @Description Retrieve a specific restaurant
// @Tags        restaurants
// @Security BearerAuth
// @Produce     json
// @Param       id path string true "Restaurant ID"
// @Success     200 {object} response.RestaurantResponse
// @Failure     404 {object} response.ErrorResponse
// @Router      /restaurant/{id} [get]
func (h *RestaurantHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.client.GetByID(context.Background(), &restaurantpb.GetRestaurantByIDRequest{Id: id})
	if err != nil {
		response.NewError(c, http.StatusNotFound, "restaurant not found")
		return
	}

	c.JSON(http.StatusOK, response.FromRestaurantProto(res))
}

// UpdateRestaurant godoc
// @Summary     Update a restaurant
// @Description Update restaurant details
// @Tags        restaurants
// @Security BearerAuth
// @Accept      json
// @Param       id path string true "Restaurant ID"
// @Param       request body restaurantpb.UpdateRestaurantRequest true "Updated info"
// @Success     200
// @Failure     400 {object} response.ErrorResponse
// @Failure     500 {object} response.ErrorResponse
// @Router      /restaurant/{id} [put]
func (h *RestaurantHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req restaurantpb.UpdateRestaurantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewError(c, http.StatusBadRequest, "invalid request body")

		return
	}
	req.Id = id

	_, err := h.client.Update(c.Request.Context(), &req)
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, "failed to update restaurant")
		return
	}

	c.Status(http.StatusOK)
}

// DeleteRestaurant godoc
// @Summary     Delete a restaurant
// @Tags        restaurants
// @Security BearerAuth
// @Param       id path string true "Restaurant ID"
// @Success     200
// @Failure     500 {object} response.ErrorResponse
// @Router      /restaurant/{id} [delete]
func (h *RestaurantHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	_, err := h.client.Delete(c.Request.Context(), &restaurantpb.DeleteRestaurantRequest{Id: id})
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, "failed to delete restaurant")
		return
	}

	c.Status(http.StatusOK)
}

// ListRestaurants godoc
// @Summary     List all restaurants
// @Tags        restaurants
// @Security BearerAuth
// @Produce     json
// @Success     200 {array} response.RestaurantResponse
// @Failure     500 {object} response.ErrorResponse
// @Router      /restaurant/all [get]
func (h *RestaurantHandler) List(c *gin.Context) {
	res, err := h.client.List(c.Request.Context(), &emptypb.Empty{})
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, "failed to list restaurants")
		return
	}

	var result []response.RestaurantResponse
	for _, r := range res.Restaurants {
		result = append(result, response.FromRestaurantProto(r))
	}
	c.JSON(http.StatusOK, result)
}
