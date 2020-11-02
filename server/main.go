package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// database connection
	OpenConnection()
	// init app
	app := gin.Default()
	// grouping routes
	v1 := app.Group("/api/v1")
	{
		// Reset & Seed Data
		v1.GET("/reset", ResetDatabaseHandler)
		v1.GET("/seed", SeedDataHandler)
		// User Actions
		v1.POST("/login", LoginUserHandler)
		v1.POST("/logout", TokenAuthMiddleware(), LogoutUserHandler)
		// GraphQL Integration
		v1.POST("/query", GraphQlMiddleware())
	}
	// Customer Actions
	customers := v1.Group("/customer")
	{
		customers.POST("/create", CreateCustomerHandler)
		customers.POST("/login", LoginCustomerHandler)
		customers.POST("/logout", TokenAuthMiddleware(), LogoutCustomerHandler)
	}
	app.Run()
}
