package main

import "fmt"

func AddStudent() {
	var newStudent Student
	var numberOfGrades int
	
	fmt.Print("Student Name: ")
	fmt.Scanln(&newStudent.Name)
	fmt.Print("Student ID: ")
	fmt.Scanln(&newStudent.Id)
	fmt.Print("Number of Grades: ")
	fmt.Scanln(&numberOfGrades)

	grades := make([]float64, numberOfGrades)
	for i := 0; i < numberOfGrades; i++ {
		fmt.Print("Grade ", (i + 1), ": ")
		fmt.Scanln(&grades[i])
	}	
	newStudent.Grades = grades

	studentLenBefore := len(students)
	students = append(students, newStudent)
	
	if len(students) > studentLenBefore {
		fmt.Println("One new student has been added!")
	}

	Landing()

}