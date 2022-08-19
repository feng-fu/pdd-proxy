package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "config",
	})
}
