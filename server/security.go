package main

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strings"
	"time"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type claims struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	jwt.StandardClaims
}

func createCustomerJwtToken(customer Customer) (string, error) {
	var err error
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &claims{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func createJwtToken(user User) (string, error) {
	var err error
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &claims{
		ID:       user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		token := strArr[1]
		// validate token on database
		user := User{}
		err := DB.Where(&User{Token: token}).First(&user).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ""
		}
		return token
	}
	return ""
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)
	claims := &claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
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

// HashPassword encrypt user password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword validate user password
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
