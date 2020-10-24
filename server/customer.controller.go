package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type customerRequest struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
}

// LoginCustomerHandler action
func LoginCustomerHandler(c *gin.Context) {
	var customerRequest customerRequest
	c.BindJSON(&customerRequest)
	customer := AuthenticationCustomerService(customerRequest.Email, customerRequest.Password)
	if customer.ID == 0 && customer.Token == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Usuário ou Senha Inválido!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": customer.Token,
		"id":    customer.ID,
		"email": &customer.Email,
		"user":  customer.FirstName + " " + customer.LastName,
	})
}

// LogoutCustomerHandler action
func LogoutCustomerHandler(c *gin.Context) {
	var customerRequest customerRequest
	c.BindJSON(&customerRequest)
	LogoutService(customerRequest.ID, customerRequest.Email)
	c.JSON(http.StatusOK, gin.H{"message": "logout..."})
}

// CreateCustomerHandler action
func CreateCustomerHandler(c *gin.Context) {
	var customerRequest customerRequest
	c.BindJSON(&customerRequest)
	customer, error := CreateCustomerService(customerRequest)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": customer.Token,
		"id":    customer.ID,
		"email": &customer.Email,
		"user":  customer.FirstName + " " + customer.LastName,
	})
}
