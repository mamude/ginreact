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
		v1.GET("/seed", SeedDataHandler)
		v1.POST("/login", LoginUserHandler)

		// companies
		v1.GET("/companies", TokenAuthMiddleware(), ListCompaniesHandler)
	}
	// starts Gin
	router.Run()
}
