package main

import (
	"fmt"
	"sort"
)

var sortedStudent []Student

func ViewStudents() {

	studentsLen := len(students)

	if studentsLen == 0 {
		fmt.Println("No students added yet")
	} else {
		if studentsLen > 1 {
			sort.Slice(students, func(i, j int) bool {
				return CalcAverageView(students[i].Grades) > 
						CalcAverageView(students[j].Grades)
			})
		}
		for i := 0; i < len(students); i++ {
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
			}
			avgOfGrades = CalcAverageView(students[i].Grades)
			if avgOfGrades >= 60 {
				status = "PASS"
			} else {
				status = "FAIL"
			}
			fmt.Println("")
			fmt.Println("Average		:", avgOfGrades)
			fmt.Println("Status		:", status)
			fmt.Println("--------------------------------")
		}
	}

	Landing()

}

func CalcAverageView(grades []float64) float64 {
	var sum float64
	for i := 0; i < len(grades); i++ {
		sum += grades[i]
	}
	return sum / float64(len(grades)) 
}