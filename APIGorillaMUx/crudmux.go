package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Dummy Database
var books []Book

func handleRequest() {
	r := mux.NewRouter()

	//Dummy Data
	books = append(books, Book{ID: "1", Isbn: "234567", Title: "Hello1", Author: Author{Firstname: "Himanshu", Lastname: "Singh"}})
	books = append(books, Book{ID: "1", Isbn: "234567", Title: "Hello2", Author: Author{Firstname: "Himanshu", Lastname: "Singh"}})
	books = append(books, Book{ID: "2", Isbn: "2345679", Title: "Hi", Author: Author{Firstname: "Himanshu", Lastname: "Singh"}})
	books = append(books, Book{ID: "3", Isbn: "2345678", Title: "Bye", Author: Author{Firstname: "Himanshu", Lastname: "Singh"}})
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", r))
}
func main() {
	handleRequest()
}
func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	var data []Book

	param := mux.Vars(r) //This is for retriving the params in this
	for _, item := range books {
		if param["id"] == item.ID {
			data = append(data, item)

		}
	}
	fmt.Println(data)
	json.NewEncoder(w).Encode(data)

}
func createBook(w http.ResponseWriter, r *http.Request) {
	var data Book
	data.ID = strconv.Itoa(rand.Intn(10000))
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("There is error occured")
		return
	}
	books = append(books, data)
	json.NewEncoder(w).Encode(books)

}
func updateBook(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	var temp Book
	for key, value := range books {
		if param["id"] == value.ID {
			temp = value
			books = append(books[:key], books[key+1:]...)

		}
	}
	var data Book
	data.ID = temp.ID
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("There is error occured")
		return
	}
	books = append(books, data)
	json.NewEncoder(w).Encode(books)

}
func deleteBook(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	for key, value := range books {
		if param["id"] == value.ID {
			books = append(books[:key], books[key+1:]...)

		}
	}

}
