package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type MenuItem struct {
	ID 		int	
	Name 	string 
	Price 	float64 
}

type Order struct {
	OrderID string 
	CustomerName string 
	Items []OrderItem 
}

type OrderItem struct {
	Menu MenuItem
	Quantity int 
}

var Menus = []MenuItem {
	{ID: 1, Name: "Espresso", Price: 25000},
	{2, "Latte", 35000},
	{3, "Cappuccino", 32000},
	{4, "Americano", 28000},
	{5, "Mocha", 38000},
	{6, "Matcha Latte", 40000},
	{7, "Croissant", 22000},
	{8, "Cheesecake", 30000},
}

var Orders []Order

func main() {
	LoadOrders()
	Landing()
}

func marshall() {
	jsonData, err := json.MarshalIndent(Orders, "", " ")

	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	err = os.WriteFile("orders.json", jsonData, 0644)
	if err != nil {
		fmt.Println("File write error :", err)
		return
	}

}

func SaveOrders() {
	marshall()
	fmt.Println("Successfully exported Orders to JSON")
}

func LoadOrders() {
	_, err := os.Stat("orders.json")

	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("File '%s' not found\n", "orders.json")
	} else if err != nil {
		fmt.Println("Error checking fil;e status:", err)
	}  else {
		fmt.Println("Loading data...")

		fileData, err := os.ReadFile("orders.json")
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		err = json.Unmarshal(fileData, &Orders)
		if err != nil {
			fmt.Println("Error parsing JSON data:", err)
			return
		}
	}

	fmt.Println("Loaded", len(Orders),"users.")
}