package main

import (
	"github.com/gin-gonic/gin"

	"github.com/mamude/ginreact/controllers"
	"github.com/mamude/ginreact/models"
)

func main() {
	// database connection
	models.OpenConnection()
	// routes
	router := gin.Default()
	// grouping routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/seed", controllers.SeedData)
		v1.POST("/login", controllers.LoginUser)

		// companies
		v1.GET("/companies", controllers.ListCompanies)
	}
	// starts Gin
	router.Run()
}
