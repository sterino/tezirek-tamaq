package routes

import (
	"api-gateway/internal/handler"
	"api-gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	router *gin.RouterGroup,
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
	restaurantHandler *handler.RestaurantHandler,
	orderHandler *handler.OrderHandler,
) {

	restaurant := router.Group("/restaurant").Use(middleware.JWTMiddleware())
	{
		restaurant.POST("/create", restaurantHandler.Create)
		restaurant.GET("/:id", restaurantHandler.GetByID)
		restaurant.PUT("/:id", restaurantHandler.Update)
		restaurant.DELETE("/:id", restaurantHandler.Delete)
		restaurant.GET("/all", restaurantHandler.List)
	}

	order := router.Group("/order").Use(middleware.JWTMiddleware())
	{
		order.POST("/create", orderHandler.Create)
		order.GET("/:id", orderHandler.GetByID)
		order.PUT("/:id/status", orderHandler.UpdateStatus)
		order.DELETE("/:id/delete", orderHandler.Delete)
		order.GET("/all", orderHandler.GetAll)
		order.GET("/user/:id", orderHandler.GetByUserID)
		order.GET("/restaurant/:id", orderHandler.GetByRestaurantID)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", authHandler.Register)
	}

	user := router.Group("/user").Use(middleware.JWTMiddleware())
	{
		user.GET("/:id", userHandler.GetByID)
		user.GET("/all", userHandler.List)
	}

}
