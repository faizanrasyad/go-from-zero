package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func CreateOrder() {
	var newOrder Order
	var orderedItems []OrderItem

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

			orderedItems = append(orderedItems, item)
		} else {
			i = -1
		}
	}

	newOrder.Items = orderedItems
	newOrder.OrderID = generateOrderID()
	Orders = append(Orders, newOrder)

	fmt.Println("")
	CreateReceipt(newOrder)
	Landing()

}

func generateOrderID() string {
	var orderID string

	if len(Orders) != 0 {
		lastOrderID := Orders[len(Orders) - 1].OrderID
		lastNum, _ := strconv.Atoi(strings.Split(lastOrderID, "D")[1])
		currentNum := lastNum + 1

		if currentNum < 10 {
			orderID = "ORD00" + strconv.Itoa(currentNum)
		} else if currentNum < 100 {
			orderID = "ORD0" + strconv.Itoa(currentNum)
		} else {
			orderID = "ORD" + strconv.Itoa(currentNum)
		}
	} else {
		orderID = "ORD001"
	}

	return orderID
}

func CreateReceipt(order Order) {
	fmt.Println("===========================")
	fmt.Println("BrewHub Coffee Cashier")
	fmt.Println("===========================")
	fmt.Println("")
	fmt.Println("Customer:", order.CustomerName)
	fmt.Println("")

	var subtotal float64
	orderedItems := order.Items
	for i := 0; i < len(orderedItems); i++ {
		priceAfterQuantity := orderedItems[i].Menu.Price * float64(orderedItems[i].Quantity)
		subtotal += priceAfterQuantity
		fmt.Println(
			orderedItems[i].Quantity,
			"x",
			orderedItems[i].Menu.Name,
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

func DeleteOrder() {
	order := SearchOrder()
	var orderIndex int

	for i := 0; i < len(Orders); i++ {
		if order.OrderID == Orders[i].OrderID {
			orderIndex = i
		}
	}

	var isSure string
	fmt.Print("Are you sure to delete this? (y/n) ")
	fmt.Scanln(&isSure)

	if isSure == "y" {
		Orders = slices.Delete(Orders, orderIndex, orderIndex+1)
		fmt.Println(order.OrderID, "order has been deleted.")
	}

	Landing()
}
