package main

import "fmt"

func AddStudent() {
	var newStudent Student
	var numberOfGrades int

	fmt.Print("Student Name: ")
	fmt.Scanln(&newStudent.Name)
	fmt.Print("Student ID: ")
	fmt.Scanln(&newStudent.Id)

	isDuplicate := false
	for i := 0; i < len(students); i++ {
		if students[i].Id == newStudent.Id {
			isDuplicate = true
		}
	}

	if isDuplicate {
		fmt.Println(
			"There is already a student with a Student ID of",
			newStudent.Id,
		)
		Landing()
		return
	}

	fmt.Print("Number of Grades: ")
	fmt.Scanln(&numberOfGrades)

	grades := make([]float64, numberOfGrades)
	for i := 0; i < numberOfGrades; i++ {
		fmt.Print("Grade ", (i + 1), ": ")
		fmt.Scanln(&grades[i])
	}
	newStudent.Grades = grades

	students = append(students, newStudent)

	Landing()

}
