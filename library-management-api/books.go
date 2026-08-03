package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func BookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetBooks(w, r)
	case http.MethodPost:
		AddBook(w, r)
	}
}

func BookByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/book/"))

	switch r.Method {
	case http.MethodGet:
		GetBookByID(w, r, id)
	case http.MethodPut:
		UpdateBook(w, r, id)
	case http.MethodDelete:
		DeleteBook(w, r, id)
	}
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Books)
}

func GetBookByID(w http.ResponseWriter, r *http.Request, id int) {
	w.Header().Set("Content-Type", "application/json")

	for _, book := range Books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	book.ID = Books[len(Books) - 1].ID + 1
	Books = append(Books, book)

	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusCreated)
}

func UpdateBook(w http.ResponseWriter, r *http.Request, id int) {
	w.Header().Set("Content-Type", "application/json")

	var updatedBook Book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i, book := range Books {
		if (book.ID == id) {
			updatedBook.ID = id
			Books[i] = updatedBook

			json.NewEncoder(w).Encode(Books[i])
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

func DeleteBook(w http.ResponseWriter, r *http.Request, id int) {
	for i, book := range Books {
		if book.ID == id {
			Books = append(Books[:i], Books[i+1:]...)

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}
