package main

import (
	"fmt"
	"strconv"
)

func CreateOrder() {
	var newOrder Order
	var orderItems []OrderItem

	fmt.Print("Customer Name: ")
	fmt.Scanln(&newOrder.CustomerName)

	for i := 1; i > 0; i++ {
		fmt.Println("")

		var item OrderItem
		var menuID int
		fmt.Print("Menu ID: ")
		fmt.Scanln(&menuID)

		if menuID != 0 {
			fmt.Print("Quantity: ")
			fmt.Scanln(&item.Quantity)
			item.Menu = Menus[menuID-1]

			orderItems = append(orderItems, item)
		} else {
			i = -1
		}
	}

	newOrder.Items = orderItems
	newOrder.OrderID = generateOrderID()
	Orders = append(Orders, newOrder)

	fmt.Println("")
	createReceipt(newOrder.CustomerName, orderItems)
	main()

}

func generateOrderID() string {
	orderLen := len(Orders)
	var orderID string

	if orderLen < 10 {
		orderID = "ORD00" + strconv.Itoa(orderLen)
	} else if orderLen < 100 {
		orderID = "ORD0" + strconv.Itoa(orderLen)
	} else {
		orderID = "ORD" + strconv.Itoa(orderLen)
	}

	return orderID
}

func createReceipt(customerName string, orderItems []OrderItem) {
	fmt.Println("===========================")
	fmt.Println("BrewHub Coffee Cashier")
	fmt.Println("===========================")
	fmt.Println("")
	fmt.Println("Customer:", customerName)
	fmt.Println("")
	
	var subtotal float64
	for i := 0; i < len(orderItems); i++ {
		priceAfterQuantity := orderItems[i].Menu.Price * float64(orderItems[i].Quantity)
		subtotal += priceAfterQuantity
		fmt.Println(
			orderItems[i].Quantity, 
			"x", 
			orderItems[i].Menu.Name,
			"	Rp",
			priceAfterQuantity,
		)
	}
	tax := subtotal * 11 / 100
	total := subtotal + tax

	fmt.Println("")
	fmt.Println("---------------------------")
	fmt.Println("")
	fmt.Println("Subtotal	Rp", subtotal)
	fmt.Println("")
	fmt.Println("Tax (11%)	Rp", tax)
	fmt.Println("")
	fmt.Println("Total		Rp", total)
	fmt.Println("")
	fmt.Println("Thank you!")
	fmt.Println("")
}
