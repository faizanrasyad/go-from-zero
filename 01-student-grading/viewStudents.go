package main

import "fmt"

func ViewStudents() {

	for i := 0; i < len(students); i++ {
		var sumOfGrades float64
		var avgOfGrades float64
		var gradeLen int = len(students[i].Grades)
		var status string

		fmt.Println("--------------------------------")
		fmt.Println("Name		:", students[i].Name)
		fmt.Println("Student ID	:", students[i].Id)
		fmt.Println("")
		fmt.Println("Grades:")
		for j := 0; j < gradeLen; j++ {
			grade := students[i].Grades[j]
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