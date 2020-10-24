package main

// AuthenticationService service
func AuthenticationService(username, password string) User {
	user := User{}
	DB.Where(&User{Username: username}).First(&user)
	match := CheckPassword(password, user.Password)
	if match {
		token, _ := createJwtToken(user)
		user.Token = token
		DB.Save(&user)
		return user
	}
	return User{}
}

// LogoutService service
func LogoutService(id int, username string) {
	user := User{}
	DB.Where(&User{ID: id, Username: username}).First(&user)
	user.Token = ""
	DB.Save(&user)
}
