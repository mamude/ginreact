package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mamude/ginreact/services"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginUser(c *gin.Context) {
	var loginRequest LoginRequest
	c.BindJSON(&loginRequest)
	user := services.Authentication(loginRequest.Username, loginRequest.Password)
	if user.ID == 0 && user.Token == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid Username or Password!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": user.Token, "username": user.Username})
}
