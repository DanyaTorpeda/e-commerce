package handlers

import (
	"e-commerce/internal/domains/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) register(c *gin.Context) {
	var input user.User

	if err := c.BindJSON(&input); err != nil {
		h.logger.Warn("invalid user data")
		NewErrorResponse(c, http.StatusBadRequest, "invalid user data")
		return
	}

}

func (h *Handler) login(c *gin.Context) {

}
