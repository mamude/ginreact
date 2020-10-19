package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// TokenAuthMiddleware checks if JWT is valid
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValidService(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
