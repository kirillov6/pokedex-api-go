package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kirillov6/pokedex-api-go/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{serv}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	pokedex := r.Group("/pokedex", h.tokenAuthMiddleware)
	{
		pokedex.GET("/", h.getPokedex)
		pokedex.POST("/", h.addPokemon)

		pokemon := pokedex.Group("/:id")
		{
			pokemon.GET("/", h.getPokemon)
			pokemon.PUT("/", h.updatePokemon)
			pokemon.DELETE("/", h.deletePokemon)
		}
	}

	return r
}
