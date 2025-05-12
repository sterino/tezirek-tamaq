package handler

import (
	"api-gateway/pkg/response"
	"api-gateway/proto/gen/orderpb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

type OrderHandler struct {
	client orderpb.OrderServiceClient
}

func NewOrderHandler(client orderpb.OrderServiceClient) *OrderHandler {
	return &OrderHandler{client: client}
}

// CreateOrder godoc
// @Summary     Create a new order
// @Description Create and return a new order
// @Tags        orders
// @Security    BearerAuth
// @Accept      json
// @Produce     json
// @Param       order body orderpb.CreateOrderRequest true "Order data"
// @Success     201 {object} response.OrderResponse
// @Failure     400 {object} response.ErrorResponse
// @Failure     500 {object} response.ErrorResponse
// @Router      /order/create [post]
func (h *OrderHandler) Create(c *gin.Context) {
	var req orderpb.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewError(c, http.StatusBadRequest, "invalid request")
		return
	}

	res, err := h.client.Create(c.Request.Context(), &req)
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, "failed to create order")
		return
	}

	c.JSON(http.StatusCreated, response.FromOrderProto(res))
}

// GetOrderByID godoc
// @Summary     Get order by ID
// @Tags        orders
// @Security    BearerAuth
// @Produce     json
// @Param       id path string true "Order ID"
// @Success     200 {object} response.OrderResponse
// @Failure     404 {object} response.ErrorResponse
// @Router      /order/{id} [get]
func (h *OrderHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.client.GetByID(context.Background(), &orderpb.GetOrderByIDRequest{Id: id})
	if err != nil {
		response.NewError(c, http.StatusNotFound, "order not found")
		return
	}

	c.JSON(http.StatusOK, response.FromOrderProto(res))
}

// UpdateOrderStatus godoc
// @Summary     Update order status
// @Tags        orders
// @Security    BearerAuth
// @Accept      json
// @Param       id path string true "Order ID"
// @Param       request body orderpb.UpdateOrderRequest true "Update data"
// @Success     200
// @Failure     400 {object} response.ErrorResponse
// @Failure     500 {object} response.ErrorResponse
// @Router      /order/{id}/status [put]
func (h *OrderHandler) UpdateStatus(c *gin.Context) {
	id := c.Param("id")
	var req orderpb.UpdateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewError(c, http.StatusBadRequest, "invalid input")
		return
	}
	req.Id = id

	_, err := h.client.Update(c.Request.Context(), &req)
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, "update failed")
		return
	}

	c.Status(http.StatusOK)
}

// DeleteOrder godoc
// @Summary     Delete order by ID
// @Tags        orders
// @Security    BearerAuth
// @Param       id path string true "Order ID"
// @Success     200
// @Failure     500 {object} response.ErrorResponse
// @Router      /order/{id}/delete [delete]
func (h *OrderHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	_, err := h.client.Delete(c.Request.Context(), &orderpb.DeleteOrderRequest{Id: id})
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, "delete failed")
		return
	}

	c.Status(http.StatusOK)
}

// ListOrders godoc
// @Summary     List all orders
// @Tags        orders
// @Security    BearerAuth
// @Produce     json
// @Success     200 {array} response.OrderResponse
// @Failure     500 {object} response.ErrorResponse
// @Router      /order/all [get]
func (h *OrderHandler) GetAll(c *gin.Context) {
	res, err := h.client.List(c.Request.Context(), &emptypb.Empty{})
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, "failed to list orders")
		return
	}

	var result []response.OrderResponse
	for _, o := range res.Orders {
		result = append(result, response.FromOrderProto(o))
	}

	c.JSON(http.StatusOK, result)
}

// ListOrdersByUser godoc
// @Summary     List orders by user ID
// @Tags        orders
// @Security    BearerAuth
// @Produce     json
// @Param       id path string true "User ID"
// @Success     200 {array} response.OrderResponse
// @Failure     500 {object} response.ErrorResponse
// @Router      /order/user/{id} [get]
func (h *OrderHandler) GetByUserID(c *gin.Context) {
	userID := c.Param("id")

	res, err := h.client.ListByUser(c.Request.Context(), &orderpb.ListByUserRequest{UserId: userID})
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, "failed to list user orders")
		return
	}

	var result []response.OrderResponse
	for _, o := range res.Orders {
		result = append(result, response.FromOrderProto(o))
	}

	c.JSON(http.StatusOK, result)
}

// ListOrdersByRestaurant godoc
// @Summary     List orders by restaurant ID
// @Tags        orders
// @Security    BearerAuth
// @Produce     json
// @Param       id path string true "Restaurant ID"
// @Success     200 {array} response.OrderResponse
// @Failure     500 {object} response.ErrorResponse
// @Router      /order/restaurant/{id} [get]
func (h *OrderHandler) GetByRestaurantID(c *gin.Context) {
	restaurantID := c.Param("id")

	res, err := h.client.ListByRestaurant(c.Request.Context(), &orderpb.ListByRestaurantRequest{RestaurantId: restaurantID})
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, "failed to list restaurant orders")
		return
	}

	var result []response.OrderResponse
	for _, o := range res.Orders {
		result = append(result, response.FromOrderProto(o))
	}

	c.JSON(http.StatusOK, result)
}
