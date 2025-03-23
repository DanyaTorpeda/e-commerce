package handlers

import (
	"e-commerce/internal/domains/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) register(c *gin.Context) {
	var input user.User

	if err := c.BindJSON(&input); err != nil {
		h.logger.Warnf("invalid user data: %s", err.Error())
		NewErrorResponse(c, http.StatusBadRequest, "invalid user data")
		return
	}

	id, err := h.service.CreateUser(input)
	if err != nil {
		h.logger.Warnf("error creating user: %s", err.Error())
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewResponse(c, "success", http.StatusCreated, "user successfully created", id)
}

func (h *Handler) login(c *gin.Context) {
	var input user.UserLogin

	if err := c.BindJSON(&input); err != nil {
		h.logger.Warnf("invalid user data: %s", err.Error())
		NewErrorResponse(c, http.StatusBadRequest, "invalid user data")
		return
	}

	usr, err := h.service.CheckUser(input)
	if err != nil {
		h.logger.Warnf("no such user found: %s", err.Error())
		NewErrorResponse(c, http.StatusUnauthorized, "no such user found")
		return
	}

	h.service.CreateToken(usr.ID, usr.Role)
}
