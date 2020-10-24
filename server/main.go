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
		// Customer Actions
		v1.POST("/customer/create", CreateCustomerHandler)
		v1.POST("/customer/login", LoginCustomerHandler)
		v1.POST("/customer/logout", TokenAuthMiddleware(), LogoutCustomerHandler)
		// Authentication & Logout
		v1.POST("/login", LoginUserHandler)
		v1.POST("/logout", TokenAuthMiddleware(), LogoutUserHandler)
		// GraphQL Integration
		v1.POST("/query", TokenAuthMiddleware(), GraphQlMiddleware())
	}
	// starts
	app.Run()
}
