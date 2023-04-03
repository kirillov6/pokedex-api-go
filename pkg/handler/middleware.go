package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userIdCtxKey        = "UserId"
)

func (h *Handler) tokenAuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader(authorizationHeader)
	if authHeader == "" {
		newResponseError(c, http.StatusUnauthorized, "Empty auth header")
		return
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 {
		newResponseError(c, http.StatusUnauthorized, "Invalid auth header")
		return
	}

	userId, err := h.service.ParseToken(authHeaderParts[1])
	if err != nil {
		newResponseError(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userIdCtxKey, userId)
	c.Next()
}
