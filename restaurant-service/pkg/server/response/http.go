package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Object struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Object{
		Success: true,
		Data:    data,
	})
}

func BadRequest(c *gin.Context, err error, data interface{}) {
	c.JSON(http.StatusBadRequest, Object{
		Success: false,
		Message: err.Error(),
		Data:    data,
	})
}

func NotFound(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, Object{
		Success: false,
		Message: err.Error(),
	})
}

func InternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, Object{
		Success: false,
		Message: err.Error(),
	})
}
