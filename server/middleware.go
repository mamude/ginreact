package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mamude/ginreact/services"
	"net/http"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := services.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
