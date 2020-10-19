package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// database connection
	OpenConnection()
	// routes
	router := gin.Default()
	// grouping routes
	v1 := router.Group("/api/v1")
	{
		// Seed Data
		v1.GET("/seed", SeedDataHandler)
		// Authentication
		v1.POST("/login", LoginUserHandler)
		// GraphQL Integration
		v1.POST("/query", TokenAuthMiddleware(), GraphQlMiddleware())
	}
	// starts Gin
	router.Run()
}
