package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:9000/", nil)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, string(body))

	fmt.Fprintf(w, validToken)
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Println("My Simple Client")

	handleRequest()
}
