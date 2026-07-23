package main

import (
	"fmt"
	"os"
)

type MenuItem struct {
	ID int
	Name string
	Price float64
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
var OrderItems []OrderItem

func main() {
	fmt.Println("===========================")
	fmt.Println("BrewHub Coffee Cashier")
	fmt.Println("===========================")
	fmt.Println("")
	fmt.Println("1. Show Menu")
	fmt.Println("2. Create Order")
	fmt.Println("3. View Orders")
	fmt.Println("4. Search Order")
	fmt.Println("5. Cancel Order")
	fmt.Println("6. Daily Sales")
	fmt.Println("7. Exit")
	fmt.Print("Choose: ")
	
	var chosen int
	fmt.Scanln(&chosen)
	fmt.Println("")

	switch chosen {
	case 1: ShowMenu()
	case 2: CreateOrder()
	case 3:
	case 4:
	case 5:
	case 6:
	case 7: os.Exit(0)
	}
}