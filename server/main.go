package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mamude/ginreact/controllers"
	"github.com/mamude/ginreact/models"
)

func main() {
	models.OpenConnection()
	app := gin.Default()

	v1 := app.Group("/api/v1")
	{
		v1.GET("/reset", controllers.ResetDatabaseHandler)
		v1.GET("/seed", controllers.SeedDataHandler)
		v1.POST("/login", controllers.LoginUserHandler)
		v1.POST("/logout", TokenAuthMiddleware(), controllers.LogoutUserHandler)
		v1.POST("/query", GraphQlMiddleware())
	}
	customers := v1.Group("/customer")
	{
		customers.POST("/create", controllers.CreateCustomerHandler)
		customers.POST("/login", controllers.LoginCustomerHandler)
		customers.POST("/logout", TokenAuthMiddleware(), controllers.LogoutCustomerHandler)
	}
	app.Run()
}
