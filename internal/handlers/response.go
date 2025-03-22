package handlers

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewErrorResponse(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, ErrorResponse{
		Status:  "error",
		Message: message,
		Code:    code,
	})
}
