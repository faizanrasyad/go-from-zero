package main

import (
	"fmt"
	"net/http"
)

func main() {

	Endpoints()

	err := http.ListenAndServe(":8080", nil)
	
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Server running on :8080")

}

func Endpoints() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/menu", GetMenu)
	http.HandleFunc("/menu/", GetMenuByID)
	http.HandleFunc("/orders", OrdersHandler)
	http.HandleFunc("/orders/", OrderByIDHandler)
}