package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Books Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Fisrtname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Create Function Handler
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	if err := json.NewEncoder(w).Encode(books); err != nil {
		log.Fatal(err)
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	param := mux.Vars(r)
	for _, book := range books {
		if book.ID == param["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	if err := json.NewEncoder(w).Encode(&Book{}); err != nil {
		log.Fatal(err)
	}
}

func createBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var newBook Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		log.Fatal(err)
	}
	newBook.ID = strconv.Itoa(rand.Intn(1000000)) // Mock ID
	books = append(books, newBook)
	if err := json.NewEncoder(w).Encode(newBook); err != nil {
		log.Fatal(err)
	}
}

func updateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	param := mux.Vars(r)
	var updateBooks Book
	if err := json.NewDecoder(r.Body).Decode(&updateBooks); err != nil {
		log.Fatal(err)
	}
	updateBooks.ID = param["id"]

	for idx, book := range books {
		if book.ID == param["id"] {
			books = append(books[:idx], books[idx+1:]...)
			books = append(books, updateBooks)
			json.NewEncoder(w).Encode(updateBooks)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	param := mux.Vars(r)
	for idx, book := range books {
		if book.ID == param["id"] {
			books = append(books[:idx], books[idx+1:]...)
			break
		}
	}
	if err := json.NewEncoder(w).Encode(books); err != nil {
		log.Fatal(err)
	}
}

// Initail books var as a slice Book struct
var books []Book

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock data for books
	books = append(books, Book{
		ID:    "1",
		Isbn:  "448742",
		Title: "Book one",
		Author: &Author{
			Fisrtname: "John",
			Lastname:  "Doe",
		},
	})

	books = append(books, Book{
		ID:    "2",
		Isbn:  "392093",
		Title: "Book two",
		Author: &Author{
			Fisrtname: "Robert",
			Lastname:  "Smite",
		},
	})

	// Route Handler // End Point
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
}
