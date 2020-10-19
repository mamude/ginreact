package services

import (
	"github.com/mamude/ginreact/models"
)

func ListCompanies() []models.Company {
	var companies []models.Company
	models.DB.Find(&companies)
	return companies
}
