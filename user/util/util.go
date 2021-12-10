package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func GenerateToken(email string) string {
	var key = []byte("MyS3cR3t")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["client"] = email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString(key)

	if err != nil {
		fmt.Errorf("Token Error: %s", err.Error())
	}

	return tokenString

}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Println("Error: ", err)
	}

	password = string(bytes)

	return password
}