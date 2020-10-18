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
	router.GET("/seed", controllers.SeedData)
	router.POST("/login", controllers.LoginUser)

	// starts Gin
	router.Run()
}
