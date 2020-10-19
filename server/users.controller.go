package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginRequest - post request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginUserHandler action
func LoginUserHandler(c *gin.Context) {
	var loginRequest LoginRequest
	c.BindJSON(&loginRequest)
	user := AuthenticationService(loginRequest.Username, loginRequest.Password)
	if user.ID == 0 && user.Token == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid Username or Password!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": user.Token, "username": user.Username})
}
