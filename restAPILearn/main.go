package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux" // third party package
)

type Article struct {
	Title   string `JSON:"title"`
	Desc    string `JSON:"desc"`
	Content string `JSON:"content"`
}

type Articles []Article // create type array of article

func allArticle(w http.ResponseWriter, r *http.Request) {
	article := Articles{
		Article{Title: "Test Title", Desc: "test Description", Content: "Test Content"},
	}

	json.NewEncoder(w).Encode(article) // encode arrticle to JSON for response it
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoint Hit") // send bytr stream
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)                          // by default function pass response and request
	myRouter.HandleFunc("/articles", allArticle).Methods("GET") // can specific Method that use in handle function

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequest()
}
