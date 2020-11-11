package services

import (
	"errors"
	"github.com/mamude/ginreact/models"
	"github.com/mamude/ginreact/requests"
	"github.com/mamude/ginreact/security"
)

var customer = models.Customer{}

// AuthenticationCustomerService service
func AuthenticationCustomerService(email, password, sessionID string) (models.Customer, error) {
	models.DB.Where(&models.Customer{Email: &email}).First(&customer)
	match := security.CheckPassword(password, customer.Password)
	if !match {
		return models.Customer{}, errors.New("Usuário ou Senha inválidos")
	}
	token, _ := security.CreateCustomerJwtToken(customer)
	customer.Token = token
	models.DB.Save(&customer)
	// sync shopping cart
	SyncShoppingCart(sessionID, customer.ID)
	return customer, nil
}

// LogoutCustomerService service
func LogoutCustomerService(id int, email string) {
	models.DB.Where(&models.Customer{ID: id, Email: &email}).First(&customer)
	customer.Token = ""
	models.DB.Save(&customer)
}

// CreateCustomerService service
func CreateCustomerService(customerRequest requests.CustomerRequest) (models.Customer, error) {
	password, _ := security.HashPassword(customerRequest.Password)
	customer := models.Customer{
		FirstName: customerRequest.FirstName,
		LastName:  customerRequest.LastName,
		Phone:     customerRequest.Phone,
		Email:     &customerRequest.Email,
		Password:  password,
	}
	if err := models.DB.Create(&customer).Error; err != nil {
		return models.Customer{}, errors.New("Já existe um email para esta conta")
	}
	token, _ := security.CreateCustomerJwtToken(customer)
	customer.Token = token
	models.DB.Save(&customer)
	return customer, nil
}
