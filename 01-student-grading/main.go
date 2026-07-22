package main

type Student struct {
	Name string
	Id string
	Grades []float64
}

var students []Student

func RemoveStudent(index int) {
	temp := make([]Student, 0)
	temp = append(temp, students[:index]...)
	temp = append(temp, students[index + 1:]...)
	students = temp
}

func main() {
	Landing()
}