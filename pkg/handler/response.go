package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Messasge string `json:"message"`
}

func newResponseError(c *gin.Context, statusCode int, message string) {
	log.Print(message)
	c.AbortWithStatusJSON(statusCode, ResponseError{message})
}
