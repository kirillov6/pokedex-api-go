package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kirillov6/pokedex-api-go/pkg/pokedex"
	"gorm.io/gorm"
)

func (h *Handler) getPokemons(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	pokemons, err := h.service.Pokedex.GetPokemons(userId)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": pokemons,
	})
}

func (h *Handler) addPokemon(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var pokemon pokedex.Pokemon
	if err := c.BindJSON(&pokemon); err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Pokedex.Create(userId, pokemon)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) getPokemon(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newResponseError(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	pokemon, err := h.service.Pokedex.GetPokemon(userId, uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newResponseError(c, http.StatusInternalServerError, "Pokemon not found")
		} else {
			newResponseError(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, pokemon)
}

func (h *Handler) updatePokemon(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newResponseError(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	var pokemon pokedex.Pokemon
	if err := c.BindJSON(&pokemon); err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Pokedex.UpdatePokemon(userId, uint(id), pokemon)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newResponseError(c, http.StatusInternalServerError, "Pokemon not found")
		} else {
			newResponseError(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (h *Handler) deletePokemon(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newResponseError(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	err = h.service.Pokedex.DeletePokemon(userId, uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newResponseError(c, http.StatusInternalServerError, "Pokemon not found")
		} else {
			newResponseError(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
