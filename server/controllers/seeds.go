package controllers

import (
	"github.com/brianvoe/gofakeit/v5"
	"github.com/gin-gonic/gin"
	"github.com/mamude/ginreact/models"
)

func SeedData(c *gin.Context) {
	// clear data
	models.DB.Exec("DELETE FROM person_roles;")
	models.DB.Exec("DELETE FROM users;")
	models.DB.Exec("DELETE FROM people;")
	models.DB.Exec("DELETE FROM companies;")
	models.DB.Exec("DELETE FROM roles;")

	// create Role
	roles := []models.Role{{Name: "Employee"}, {Name: "Manager"}, {Name: "CEO"}}
	models.DB.Create(&roles)

	// create User
	password := "admin"
	hash, _ := models.HashPassword(password)
	models.DB.Create(&models.User{Username: "admin", Password: hash})

	// create companies
	for i := 0; i < 100; i++ {
		name := gofakeit.Company()
		company := &models.Company{
			OpeningYear: gofakeit.Year(),
			Name:        &name,
			Description: gofakeit.BuzzWord(),
		}
		models.DB.Create(company)
	}

	// create people
	for i := 0; i < 500; i++ {
		company := models.Company{}
		models.DB.First(&company, gofakeit.Number(1, 100))

		person := &models.Person{
			Age:       gofakeit.Number(18, 65),
			Level:     gofakeit.JobTitle(),
			Country:   gofakeit.Country(),
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Salary:    gofakeit.Price(1.000, 10.000),
			Company:   company,
			CompanyID: company.ID,
		}
		models.DB.Create(person)
	}

	// associate person x role
	for i := 0; i < 200; i++ {
		// get role
		role := models.Role{}
		models.DB.First(&role, gofakeit.Number(1, 3))

		// get person
		person := models.Person{}
		models.DB.First(&person, gofakeit.Number(1, 500))
		models.DB.Model(&person).Association("Roles").Append(&role)
	}
	c.JSON(200, gin.H{"message": "Added fake data..."})
}
