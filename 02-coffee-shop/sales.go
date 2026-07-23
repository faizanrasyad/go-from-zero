package main

import (
	"fmt"
	"strconv"
)

func DailySales() {
	fmt.Println("===== Daily Sales =====")
	fmt.Println("")
	fmt.Println("Orders :", totalOrders())
	fmt.Println("")
	fmt.Println("Revenue : Rp", totalRevenue())
	fmt.Println("")
	menu, sold := bestSeller()
	fmt.Println("Best Seller :", menu, "(" + strconv.Itoa(sold), "sold)")
	fmt.Println("")
	fmt.Println("Average Order : Rp", averageOrder())
	bestSeller()
	averageOrder()
}

func averageOrder() float64 {
	var sum float64

	for i := 0; i < len(Orders); i++ {
		for j := 0; j < len(Orders[i].Items); j++ {
			sum += Orders[i].Items[j].Menu.Price * float64(Orders[i].Items[j].Quantity)
		}
	}

	return sum / float64(len(Orders))
}

func bestSeller() (menu string, sold int) {
	var itemSold = make([]int, len(Menus))
	for i := 0; i < len(Orders); i++ {
		for j := 0; j < len(Orders[i].Items); j++ {
			itemSold[Orders[i].Items[j].Menu.ID - 1] += Orders[i].Items[j].Quantity
		}
	}

	tempMax := 0
	var indexMax int
	for i := 0; i < len(itemSold); i++ {
		if (itemSold[i] > tempMax) {
			tempMax = itemSold[i]
			indexMax = i
		}
	}

	return Menus[indexMax].Name, itemSold[indexMax]
}

func totalRevenue() float64 {
	var sum float64

	for i := 0; i < len(Orders); i++ {
		for j := 0; j < len(Orders[i].Items); j++ {
			sum += Orders[i].Items[j].Menu.Price * float64(Orders[i].Items[j].Quantity)
		}
	}

	return sum
}

func totalOrders() int {
	return len(Orders)
}