package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// database connection
	OpenConnection()
	// routes
	app := gin.Default()
	// grouping routes
	v1 := app.Group("/api/v1")
	{
		// Seed Data
		v1.GET("/seed", SeedDataHandler)
		// Authentication
		v1.POST("/login", LoginUserHandler)
		// GraphQL Integration
		v1.POST("/query", TokenAuthMiddleware(), GraphQlMiddleware())
	}
	// starts Gin
	app.Run()
}
