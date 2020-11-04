package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mamude/ginreact/requests"
	"github.com/mamude/ginreact/services"
	"net/http"
)

var customerRequest = requests.CustomerRequest{}

// LoginCustomerHandler action
func LoginCustomerHandler(c *gin.Context) {
	c.BindJSON(&customerRequest)
	customer, err := services.AuthenticationCustomerService(customerRequest.Email, customerRequest.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	c.BindJSON(&customerRequest)
	services.LogoutService(customerRequest.ID, customerRequest.Email)
	c.JSON(http.StatusOK, gin.H{"message": "logout..."})
}

// CreateCustomerHandler action
func CreateCustomerHandler(c *gin.Context) {
	c.BindJSON(&customerRequest)
	customer, err := services.CreateCustomerService(customerRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": customer.Token,
		"id":    customer.ID,
		"email": &customer.Email,
		"user":  customer.FirstName + " " + customer.LastName,
	})
}
