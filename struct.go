package main

import "fmt"

type Student struct {
	NIM, Name, Surname string
	Age                int
	Gpa                float64
}
type User struct {
	Age      int
	ID, Name string
	Gpa      float64
}

func main() {

	var student Student
	student.NIM = "2210511023"
	student.Name = "Hary Capri"
	student.Surname = "Hary"
	student.Age = 21
	student.Gpa = 3.92

	// Struct Literals
	studentOb := &User{
		ID:   student.NIM,
		Name: student.Name,
		Age:  student.Age,
		Gpa:  student.Gpa,
	}

	studentDua := &User{19, "2210511023", "Hary Capri", 3.92}

	fmt.Println(student)
	fmt.Println(studentOb)
	fmt.Println(studentDua)
}
