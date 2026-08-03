package main

import (
	"fmt"
	"net/http"
)

func main() {
	Endpoints()

	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Server running on :8081")
}

func Endpoints() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/book", BookHandler)
	http.HandleFunc("/book/", BookByIdHandler)
}

func Home(w http.ResponseWriter, r*http.Request) {
	fmt.Fprintln(w, "Welcome to Library Management API!")
}