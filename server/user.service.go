package main

import "errors"

// AuthenticationService service
func AuthenticationService(username, password string) (User, error) {
	user := User{}
	DB.Where(&User{Username: username}).First(&user)
	match := CheckPassword(password, user.Password)
	if !match {
		return User{}, errors.New("Usuário ou Senha inválidos")
	}
	token, _ := createJwtToken(user)
	user.Token = token
	DB.Save(&user)
	return user, nil
}

// LogoutService service
func LogoutService(id int, username string) {
	user := User{}
	DB.Where(&User{ID: id, Username: username}).First(&user)
	user.Token = ""
	DB.Save(&user)
}
