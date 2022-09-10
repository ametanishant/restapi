package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Author Struct
type Author struct {
	FirstName string `json:"firstanme"`
	LastName  string `json:"lastname"`
}

//Book Struct(Model)
type Book struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}

//Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Something bad happenedd!"))

}

//Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {

}

//Create New Book
func createBook(w http.ResponseWriter, r *http.Request) {

}

//Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

//Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//Init Router
	r := mux.NewRouter()

	//Route handlers /Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id]", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id]", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
