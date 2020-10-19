package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mamude/ginreact/models"
	"net/http"
	"os"
	"strings"
	"time"
)

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func createJwtToken(user models.User) (string, error) {
	var err error
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		ID:       user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(JwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return token, err
		}
		return token, err
	}
	if !token.Valid {
		return token, err
	}
	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := verifyToken(r)
	if err != nil {
		return err
	}
	if !token.Valid {
		return err
	}
	return nil
}

func Authentication(username, password string) models.User {
	user := models.User{}
	models.DB.Where("username = ?", username).First(&user)
	match := models.CheckPassword(password, user.Password)
	if match {
		token, _ := createJwtToken(user)
		user.Token = token
		models.DB.Save(&user)
		return user
	}
	return models.User{}
}
