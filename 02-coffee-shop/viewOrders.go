package main

import "fmt"

func ViewOrders() {
	for i := 0; i < len(Orders); i++ {
		fmt.Println("--------------------------------")
		fmt.Println("Order ID :", Orders[i].OrderID)
		fmt.Println("")
		fmt.Println("Customer :", Orders[i].CustomerName)
		fmt.Println("")
		fmt.Println("Items")
		fmt.Println("")

		orderedItems := Orders[i].Items
		var subtotal float64
		var tax float64
		var total float64
		for j := 0; j < len(orderedItems); j++ {
			priceAfterQuantity := orderedItems[j].Menu.Price * float64(orderedItems[j].Quantity)
			subtotal += priceAfterQuantity
			fmt.Println(orderedItems[j].Quantity, orderedItems[j].Menu.Name)
			fmt.Println("")
		}
		tax = subtotal * 11 / 100
		total = subtotal + tax

		fmt.Println("Total :	Rp", total)
		fmt.Println("--------------------------------")
	}

	Landing()
}