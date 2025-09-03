package main

import "fmt"

type Subjects struct {
	Title string
	Sks   int
}

type Students struct {
	Subjects []Subjects
	Name     string
	Nim      string
	Gpa      float64
}

func (student Students) semester(name int) {
	fmt.Println("Your Subject in semester ", name)
	for i, sub := range student.Subjects {
		fmt.Printf("%d) %s (%d SKS)\n", i+1, sub.Title, sub.Sks)
	}
	fmt.Println(student.Name)
	fmt.Println(student.Nim)
	fmt.Println(student.Gpa)
}

func main() {
	student := Students{
		Subjects: []Subjects{
			{Title: "MTK DISKRIT", Sks: 3},
		},
		Name: "Michael",
		Nim:  "2210511023",
		Gpa:  3.92}

	student.semester(1)
}
