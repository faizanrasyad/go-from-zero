package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Home (w http.ResponseWriter, r*http.Request) {
	fmt.Fprintln(w, "Welcome to BrewHub Coffee API!")
}

func GetMenu (w http.ResponseWriter, r*http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Menus)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func OrdersHandler (w http.ResponseWriter, r*http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetOrders(w, r)
	case http.MethodPost:
		CreateOrder(w, r)
	}
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	// Do Create Order!
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Orders)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}