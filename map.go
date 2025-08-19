package main

import "fmt"

func main() {
	//var student map[string]string = make(map[string]string)
	//student["name"] = "Hary"
	//student["age"] = "20"

	var students = map[string]string{
		"name":     "Hary Capri",
		"age":      "18",
		"semester": "7",
		"major":    "Informatics",
	}

	fmt.Println(students["name"])
	fmt.Println(students["age"])
	fmt.Println(students["semester"])
	fmt.Println(students["major"])
	fmt.Println(students["majorx"])
	fmt.Println(students)

	// Function Map
	fmt.Println(len(students))
	fmt.Println(students["major"])
	students["majorx"] = "InformaticsX"
	fmt.Println(students["majorx"])

	var studentCopy = make(map[string]string)
	studentCopy = students
	fmt.Println(studentCopy)
	delete(studentCopy, "name")
	fmt.Println("Hapus name : ", studentCopy["name"])

}
