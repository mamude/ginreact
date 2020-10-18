package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mamude/ginreact/models"
	"os"
	"time"
)

func createJwtToken(userID uint) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Hour * 5).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func Authentication(username, password string) models.User {
	user := models.User{}
	models.DB.Where("username = ?", username).First(&user)
	match := models.CheckPassword(password, user.Password)
	if match {
		// update user token
		token, _ := createJwtToken(user.ID)
		user.Token = token
		models.DB.Save(&user)
		return user
	}
	return models.User{}
}
