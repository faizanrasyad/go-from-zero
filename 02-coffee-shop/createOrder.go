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

	fmt.Println("New order has been added")
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
