package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kirillov6/pokedex-api-go/pkg/pokedex"
)

func (h *Handler) signUp(c *gin.Context) {
	var user pokedex.User
	if err := c.BindJSON(&user); err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.CreateUser(user)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var creds pokedex.Creditionals
	if err := c.BindJSON(&creds); err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.GenerateToken(creds)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
