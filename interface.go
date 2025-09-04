package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHellow(value HasName) {
	fmt.Println("Hello my name,", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	sayHellow(Person{Name: "Hary"})

}
