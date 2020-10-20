package main

import (
	"github.com/brianvoe/gofakeit/v5"
	"github.com/gin-gonic/gin"
	"strings"
)

func resetTables() {
	DB.Exec("DELETE FROM users;")
	DB.Exec("DELETE FROM people;")
	DB.Exec("DELETE FROM companies;")
	DB.Exec("DELETE FROM roles;")
	DB.Exec("DELETE FROM order_items;")
	DB.Exec("DELETE FROM orders;")
	DB.Exec("DELETE FROM shopping_carts;")
	DB.Exec("DELETE FROM products;")
	DB.Exec("DELETE FROM markets;")
	DB.Exec("DELETE FROM customers;")
}

func seedTables() {
	// create User
	password := "admin"
	hash, _ := HashPassword(password)
	DB.Create(&User{Username: "admin", Password: hash})

	// create markets
	for i := 0; i < 20; i++ {
		appName := gofakeit.AppName()
		DB.Create(&Market{
			Name:        &appName,
			Description: gofakeit.Phrase(),
			Email:       gofakeit.Email(),
			URL:         gofakeit.URL(),
			Phone:       gofakeit.Phone(),
		})
	}

	// create products
	for i := 0; i < 100; i++ {
		market := Market{}
		DB.First(&market, gofakeit.Number(1, 10))
		productName := []string{
			gofakeit.BeerName(),
			gofakeit.BeerMalt(),
			gofakeit.Breakfast(),
			gofakeit.Fruit(),
			gofakeit.Vegetable(),
			gofakeit.Lunch(),
			gofakeit.Snack(),
		}
		DB.Create(&Product{
			MarketID:    market.ID,
			Name:        gofakeit.RandomString(productName),
			Description: gofakeit.LoremIpsumSentence(50),
			Price:       gofakeit.Price(1, 50),
			Amount:      gofakeit.Number(1, 100),
		})
	}

	// create customers
	for i := 0; i < 100; i++ {
		lastName := gofakeit.LastName()
		email := strings.ToLower(lastName + "@" + gofakeit.DomainName())
		password, _ = HashPassword("123456")
		DB.Create(&Customer{
			FirstName: gofakeit.FirstName(),
			LastName:  lastName,
			Phone:     gofakeit.Phone(),
			Email:     &email,
			Password:  password,
		})
	}
}

// ResetDatabaseHandler action
func ResetDatabaseHandler(c *gin.Context) {
	go resetTables()
	c.JSON(200, gin.H{"message": "Reseting database..."})
}

// SeedDataHandler action
func SeedDataHandler(c *gin.Context) {
	go seedTables()
	c.JSON(200, gin.H{"message": "Adding fake data..."})
}
