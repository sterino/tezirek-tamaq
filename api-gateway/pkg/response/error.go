package response

import "github.com/gin-gonic/gin"

// ErrorResponse представляет структуру ошибки, возвращаемую API
type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"invalid request"`
}

func NewError(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, ErrorResponse{
		Code:    statusCode,
		Message: message,
	})
}
