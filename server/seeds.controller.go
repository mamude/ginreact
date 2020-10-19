package main

import (
	"github.com/brianvoe/gofakeit/v5"
	"github.com/gin-gonic/gin"
)

// SeedDataHandler action
func SeedDataHandler(c *gin.Context) {
	// clear data
	DB.Exec("DELETE FROM person_roles;")
	DB.Exec("DELETE FROM users;")
	DB.Exec("DELETE FROM people;")
	DB.Exec("DELETE FROM companies;")
	DB.Exec("DELETE FROM roles;")

	// create Role
	roles := []Role{{Name: "Employee"}, {Name: "Manager"}, {Name: "CEO"}}
	DB.Create(&roles)

	// create User
	password := "admin"
	hash, _ := HashPassword(password)
	DB.Create(&User{Username: "admin", Password: hash})

	// create companies
	for i := 0; i < 100; i++ {
		name := gofakeit.Company()
		company := &Company{
			OpeningYear: gofakeit.Year(),
			Name:        &name,
			Description: gofakeit.BuzzWord(),
		}
		DB.Create(company)
	}

	// create people
	for i := 0; i < 500; i++ {
		company := Company{}
		DB.First(&company, gofakeit.Number(1, 100))

		person := &Person{
			Age:       gofakeit.Number(18, 65),
			Level:     gofakeit.JobTitle(),
			Country:   gofakeit.Country(),
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Salary:    gofakeit.Price(1.000, 10.000),
			Company:   company,
			CompanyID: company.ID,
		}
		DB.Create(person)
	}

	// associate person x role
	for i := 0; i < 200; i++ {
		// get role
		role := Role{}
		DB.First(&role, gofakeit.Number(1, 3))

		// get person
		person := Person{}
		DB.First(&person, gofakeit.Number(1, 500))
		DB.Model(&person).Association("Roles").Append(&role)
	}
	c.JSON(200, gin.H{"message": "Added fake data..."})
}
