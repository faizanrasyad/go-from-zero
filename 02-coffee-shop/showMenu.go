package main

import (
	"fmt"
	"strconv"
)

func ShowMenu() {
	fmt.Println("All Menus")
	for i := 0; i < len(Menus); i++ {
		count := i + 1
		fmt.Println(strconv.Itoa(count) + ".", Menus[i].Name + "		Rp", Menus[i].Price)
	}
	fmt.Println("")
	Landing()
}