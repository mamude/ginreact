package main

import "errors"

// AuthenticationCustomerService service
func AuthenticationCustomerService(email, password string) (Customer, error) {
	customer := Customer{}
	DB.Where(&Customer{Email: &email}).First(&customer)
	match := CheckPassword(password, customer.Password)
	if !match {
		return Customer{}, errors.New("Usuário ou Senha inválidos")
	}
	token, _ := createCustomerJwtToken(customer)
	customer.Token = token
	DB.Save(&customer)
	return customer, nil
}

// LogoutCustomerService service
func LogoutCustomerService(id int, email string) {
	customer := Customer{}
	DB.Where(&Customer{ID: id, Email: &email}).First(&customer)
	customer.Token = ""
	DB.Save(&customer)
}

// CreateCustomerService service
func CreateCustomerService(customerRequest customerRequest) (Customer, error) {
	password, _ := HashPassword(customerRequest.Password)
	customer := Customer{
		FirstName: customerRequest.FirstName,
		LastName:  customerRequest.LastName,
		Phone:     customerRequest.Phone,
		Email:     &customerRequest.Email,
		Password:  password,
	}
	if err := DB.Create(&customer).Error; err != nil {
		return Customer{}, errors.New("Já existe um email para esta conta")
	}
	token, _ := createCustomerJwtToken(customer)
	customer.Token = token
	DB.Save(&customer)
	return customer, nil
}
