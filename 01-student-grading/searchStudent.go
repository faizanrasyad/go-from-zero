package main

import (
	"fmt"
	"strconv"
)

func SearchStudent() {
	var studentID int
	fmt.Print("Enter Student ID: ")
	fmt.Scanln(&studentID)

	var chosenStudent Student
	for i := 0; i < len(students); i++ {
		if students[i].Id == strconv.Itoa(studentID) {
			chosenStudent = students[i]
		}
	}

	if chosenStudent.Id == "" {
		fmt.Println("Student not found.")
	} else {
		var sumOfGrades float64
		var avgOfGrades float64
		var gradeLen int = len(chosenStudent.Grades)
		var status string

		fmt.Println("--------------------------------")
		fmt.Println("Name		:", chosenStudent.Name)
		fmt.Println("Student ID	:", chosenStudent.Id)
		fmt.Println("")
		fmt.Println("Grades:")
		for j := 0; j < gradeLen; j++ {
			grade := chosenStudent.Grades[j]
			fmt.Println(grade)
			sumOfGrades += grade
		}
		avgOfGrades = sumOfGrades / float64(gradeLen)
		if avgOfGrades >= 60 { 
			status = "PASS"
		} else { 
			status = "FAIL"
		}
		fmt.Println("")
		fmt.Println("Average	:", avgOfGrades)
		fmt.Println("Status		:", status)
		fmt.Println("--------------------------------")
	}

	Landing()
}