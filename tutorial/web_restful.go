/*
DELETE http://localhost:9090/1
Accept: application/json

###
GET http://localhost:9090/
Accept: application/json

###
POST http://localhost:9090/
Content-Type: application/json

	{
	  "title": "java"
	}

###

PUT http://localhost:9090/1
Content-Type: application/json

	{
	  "title": "spring"
	}
*/
package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Book is a struct that contains information about a book.
// It contains fields for an ID and a title.
type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var books = make([]Book, 0)

func main() {
	r := chi.NewRouter()

	// https://go-chi.io/#/pages/middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Get("/", index)
	r.Get("/{id}", show)
	r.Post("/", create)
	r.Put("/{id}", update)
	r.Delete("/{id}", destroy)

	http.ListenAndServe(":9090", r)
}

func destroy(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for i, book := range books {
		if strconv.Itoa(book.ID) == id {
			books = append(books[:i], books[i+1:]...)
			_ = json.NewEncoder(w).Encode(map[string]interface{}{
				"code":    0,
				"message": "success",
			})
			return
		}
	}

	http.NotFound(w, r)
}

func update(w http.ResponseWriter, r *http.Request) {
	var updatedBook Book
	id := chi.URLParam(r, "id")

	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil || updatedBook.Title == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	for i, book := range books {
		if strconv.Itoa(book.ID) == id {
			books[i].Title = updatedBook.Title
			_ = json.NewEncoder(w).Encode(books[i])
			return
		}
	}

	http.NotFound(w, r)
}

// create adds a book to the collection
func create(w http.ResponseWriter, r *http.Request) {
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil || book.Title == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	book.ID = len(books) + 1
	books = append(books, book)
	w.Header().Set("Location", r.URL.Path+"/"+strconv.Itoa(book.ID))
	_ = json.NewEncoder(w).Encode(book)
}

// show return a specific book
func show(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for _, book := range books {
		if strconv.Itoa(book.ID) == id {
			_ = json.NewEncoder(w).Encode(book)
		}
	}
	http.NotFound(w, r)
}

// index is an http.HandlerFunc that handles requests for a list of books. It produces a response in JSON format.
func index(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(books)
}
