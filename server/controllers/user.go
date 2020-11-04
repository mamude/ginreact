package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mamude/ginreact/requests"
	"github.com/mamude/ginreact/services"
	"net/http"
)

var userLoginRequest = requests.UserLoginRequest{}

// LoginUserHandler action
func LoginUserHandler(c *gin.Context) {
	c.BindJSON(&userLoginRequest)
	user, err := services.AuthenticationService(userLoginRequest.Username, userLoginRequest.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": user.Token, "username": user.Username})
}

// LogoutUserHandler action
func LogoutUserHandler(c *gin.Context) {
	c.BindJSON(&userLoginRequest)
	services.LogoutService(userLoginRequest.ID, userLoginRequest.Username)
	c.JSON(http.StatusOK, gin.H{"message": "logout..."})
}
