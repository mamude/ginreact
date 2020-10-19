package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListCompaniesHandler action
func ListCompaniesHandler(c *gin.Context) {
	companies := ListCompaniesService()
	c.JSON(http.StatusOK, gin.H{"companies": companies})
}
