package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseOK(c *gin.Context, body interface{}) {
	c.JSON(http.StatusOK, body)
}

func ResponseBadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, Message{
		Message: message,
	})
}

func ResponseNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, Message{
		Message: message,
	})
}

func ResponseCreated(c *gin.Context, body interface{}) {
	c.JSON(http.StatusCreated, body)
}
