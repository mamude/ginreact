package services

import (
	"errors"
	"github.com/mamude/ginreact/models"
	"github.com/mamude/ginreact/security"
)

var user = models.User{}

// AuthenticationService service
func AuthenticationService(username, password string) (models.User, error) {
	models.DB.Where(&models.User{Username: username}).First(&user)
	match := security.CheckPassword(password, user.Password)
	if !match {
		return models.User{}, errors.New("Usuário ou Senha inválidos")
	}
	token, _ := security.CreateUserJwtToken(user)
	user.Token = token
	models.DB.Save(&user)
	return user, nil
}

// LogoutService service
func LogoutService(id int, username string) {
	models.DB.Where(&models.User{ID: id, Username: username}).First(&user)
	user.Token = ""
	models.DB.Save(&user)
}
