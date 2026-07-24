package main

import (
	"fmt"
	"os"
)

func Landing() {
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
	case 3: ViewOrders()
	case 4: SearchOrder()
			Landing()
	case 5: DeleteOrder()
	case 6: DailySales()
	case 7: SaveOrders()
			os.Exit(0)
	}
}