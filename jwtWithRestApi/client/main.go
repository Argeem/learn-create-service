package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("mysupersecretphrase")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256) // crate new jwt token

	claims := token.Claims.(jwt.MapClaims) // add chaims detail

	claims["authorized"] = true
	claims["user"] = "Sitthisak Bannob"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey) // sign token with secret key

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
	}
	return tokenString, err
}

func main() {
	fmt.Println("My Simple Client")

	tokenString, err := GenerateJWT()
	if err != nil {
		fmt.Println("Error generating token")
	}

	fmt.Println(tokenString)
}
