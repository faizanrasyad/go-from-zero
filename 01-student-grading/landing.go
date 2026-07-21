package main

import (
	"fmt"
	"os"
)

func Landing() {
	fmt.Println("====================")
	fmt.Println("Student Grade Management")
	fmt.Println("====================")

	fmt.Println("1. Add Student")
	fmt.Println("2. View Students")
	fmt.Println("3. Search Student")
	fmt.Println("4. Show Statistics")
	fmt.Println("5. Exit")

	var menuChosen int
	fmt.Print("Choose: ")
	fmt.Scanln(&menuChosen)

	switch menuChosen{
	case 1: AddStudent()
	case 2: ViewStudents()
	case 3: SearchStudent()
	case 4: Statistics()
	case 5: os.Exit(0)
	}
}