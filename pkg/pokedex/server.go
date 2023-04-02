package pokedex

import "github.com/gin-gonic/gin"

type Server struct {
	router *gin.Engine
}

func NewServer(r *gin.Engine) *Server {
	return &Server{r}
}

func (s *Server) Run() error {
	return s.router.Run()
}
