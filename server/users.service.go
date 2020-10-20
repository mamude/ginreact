package main

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
	"time"
)

// JwtSecret env
var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Claims jwt
type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func createJwtToken(user User) (string, error) {
	var err error
	expirationTime := time.Now().Add(24 * time.Hour)
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

// TokenValidService service
func TokenValidService(r *http.Request) error {
	token, err := verifyToken(r)
	if err != nil {
		return err
	}
	if !token.Valid {
		return err
	}
	return nil
}

// AuthenticationService service
func AuthenticationService(username, password string) User {
	user := User{}
	DB.Where("username = ?", username).First(&user)
	match := CheckPassword(password, user.Password)
	if match {
		token, _ := createJwtToken(user)
		user.Token = token
		DB.Save(&user)
		return user
	}
	return User{}
}
