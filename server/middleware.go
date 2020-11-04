package main

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
	"github.com/mamude/ginreact/gql"
	"github.com/mamude/ginreact/security"
	"net/http"
)

// TokenAuthMiddleware checks if JWT is valid
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := security.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}

// GraphQlMiddleware middleware
func GraphQlMiddleware() gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema:     &gql.EcommerceSchema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: false,
	})
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
