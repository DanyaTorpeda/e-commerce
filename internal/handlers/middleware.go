package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	tokenKey            = "Token"
)

func (h *Handler) CheckToken(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		h.logger.Warn("no header")
		NewErrorResponse(c, http.StatusUnauthorized, "no header")
		c.Abort()
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		h.logger.Warn("invalid authorization format")
		NewErrorResponse(c, http.StatusUnauthorized, "invalid authorization format. Expected 'Bearer <token>'")
		c.Abort()
		return
	}

	c.Set(tokenKey, headerParts[1])
	c.Next()
}

func (h *Handler) GetId(c *gin.Context) (int, error) {
	val, ok := c.Get(tokenKey)
	if !ok {
		h.logger.Warn("no header data")
		NewErrorResponse(c, http.StatusBadRequest, "no header")
		return 0, errors.New("no header")
	}

	token, ok := val.(string)
	if !ok {
		h.logger.Warn("invalid token data")
		NewErrorResponse(c, http.StatusBadRequest, "invalid token data")
		return 0, errors.New("invalid token data")
	}

	claims, err := h.service.ParseToken(token)
	if err != nil {
		h.logger.Warnf("error parsing token: %s", err.Error())
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return 0, err
	}

	return claims.UserID, nil
}
