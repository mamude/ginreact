package main

import (
	"strings"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/gin-gonic/gin"
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
	DB.Exec("DELETE FROM categories;")
	DB.Exec("DELETE FROM category_businesses;")
	DB.Exec("DELETE FROM customers;")
}

func seedTables() {
	// create User
	password := "admin"
	hash, _ := HashPassword(password)
	DB.Create(&User{Username: "admin", Password: hash})

	// create category business
	categoryBusiness := []CategoryBusiness{
		{Name: "Category 1"},
		{Name: "Category 2"},
		{Name: "Category 3"},
		{Name: "Category 4"},
		{Name: "Category 5"},
		{Name: "Category 6"},
		{Name: "Category 7"},
		{Name: "Category 8"},
		{Name: "Category 9"},
		{Name: "Category 10"},
	}
	DB.Create(&categoryBusiness)

	// create markets
	for i := 0; i < 20; i++ {
		categoryBusiness := CategoryBusiness{}
		DB.First(&categoryBusiness, gofakeit.Number(1, 10))
		appName := gofakeit.AppName()
		DB.Create(&Market{
			CategoryBusinessID: categoryBusiness.ID,
			Name:               &appName,
			Description:        gofakeit.Phrase(),
			Email:              gofakeit.Email(),
			URL:                gofakeit.URL(),
			Phone:              gofakeit.Phone(),
		})
	}

	// create categories
	categories := []Category{
		{Name: "Category 1"},
		{Name: "Category 2"},
		{Name: "Category 3"},
		{Name: "Category 4"},
		{Name: "Category 5"},
		{Name: "Category 6"},
		{Name: "Category 7"},
		{Name: "Category 8"},
		{Name: "Category 9"},
		{Name: "Category 10"},
	}
	DB.Create(&categories)

	// create products
	for i := 0; i < 100; i++ {
		market := Market{}
		category := Category{}
		DB.First(&category, gofakeit.Number(1, 10))
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
			CategoryID:  category.ID,
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
