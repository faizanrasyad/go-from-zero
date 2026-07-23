package main

import "fmt"

func SearchOrder() {
	var orderID string
	var order Order
	fmt.Print("Order ID : ")
	fmt.Scanln(&orderID)
	for i := 0; i < len(Orders); i++ {
		if (orderID == Orders[i].OrderID) {
			order = Orders[i]
		}
	}

	if order.OrderID != "" {
		CreateReceipt(order)
	} else {
		fmt.Println("Order not found.")
	}

	main()
	
}