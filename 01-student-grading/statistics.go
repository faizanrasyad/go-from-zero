package main

import "fmt"

func Statistics() {
	fmt.Println("===== Statistics =====")
	fmt.Println("")
	fmt.Println("Total Students : ", len(students))
	fmt.Println("")

	if len(students) != 0 {
		var smartStudentName,
			smartStudentGrade,
			dumbStudentName,
			dumbStudentGrade,
			overallAverage = calcAverage()

		fmt.Println("Highest Average")
		fmt.Println(smartStudentName, "(", smartStudentGrade, ")")
		fmt.Println("")
		fmt.Println("Lowest Average")
		fmt.Println(dumbStudentName, "(", dumbStudentGrade, ")")
		fmt.Println("")
		fmt.Println("Overall Average")
		fmt.Println(overallAverage)
	}

	Landing()

}

func calcAverage() (smartStudentName string,
	smartStudentGrade float64,
	dumbStudentName string,
	dumbStudentGrade float64,
	overallAverage float64) {

	var studentOverallGrades = []float64{}
	for i := 0; i < len(students); i++ {
		var sumOfGrades float64
		var avgOfGrades float64
		for j := 0; j < len(students[i].Grades); j++ {
			sumOfGrades += students[i].Grades[j]
		}
		avgOfGrades = sumOfGrades / float64(len(students[i].Grades))
		studentOverallGrades = append(studentOverallGrades, avgOfGrades)
	}
	var sumOverallGrades float64
	smartStudent := 0
	dumbStudent := 0
	for i := 0; i < len(studentOverallGrades); i++ {
		if studentOverallGrades[i] > studentOverallGrades[smartStudent] {
			smartStudent = i
		}
		if studentOverallGrades[i] < studentOverallGrades[dumbStudent] {
			dumbStudent = i
		}
		sumOverallGrades += studentOverallGrades[i]
	}

	overallAverage = sumOverallGrades / float64(len(studentOverallGrades))
	smartStudentName = students[smartStudent].Name
	smartStudentGrade = studentOverallGrades[smartStudent]
	dumbStudentName = students[dumbStudent].Name
	dumbStudentGrade = studentOverallGrades[dumbStudent]

	return
}
