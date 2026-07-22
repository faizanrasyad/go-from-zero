package main

import (
	"fmt"
	"strings"
)

func SearchStudent() {
	var studentID string
	var studentIndex int
	fmt.Print("Enter Student ID: ")
	fmt.Scanln(&studentID)

	var chosenStudent Student
	for i := 0; i < len(students); i++ {
		if  students[i].Id == studentID {
			chosenStudent = students[i]
			studentIndex = i
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
		fmt.Println("Average		:", avgOfGrades)
		fmt.Println("Status		:", status)
		fmt.Println("--------------------------------")
		fmt.Println("")

		var action string
		fmt.Print("Actions (edit-grades/del/esc)? ")
		fmt.Scanln(&action)

		switch action {
		case "edit-grades":  
			grades := make([]float64, gradeLen)
			for i := 0; i < gradeLen; i++ {
				fmt.Print("Grade ", (i + 1), ": ")
				fmt.Scanln(&grades[i])
			}
			students[studentIndex].Grades = grades
			fmt.Println("Grades updated successfully")
			SearchStudent()
		case "del":
			var isSure string
			fmt.Print("Are you sure to delete", chosenStudent.Name, "(y/n)?")
			fmt.Scanln(&isSure)

			if strings.EqualFold("y", isSure) {
				RemoveStudent(studentIndex)
				fmt.Println(chosenStudent.Name, "has been deleted")
			}
		default: 
		}

	}

	Landing()

}