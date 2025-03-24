package handlers

import (
	"e-commerce/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	logger  *logrus.Logger
	service *services.Service
}

func New(service *services.Service, logger *logrus.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.register)
		auth.POST("/login", h.login)
	}

	users := router.Group("/users", h.CheckToken)
	{
		users.GET("/me", h.getMe)
		users.PUT("/me", h.updateMe)
		users.GET("/:id", h.getUser)
		users.DELETE("/:id", h.deleteUser)
	}

	return router
}
