package main

import "fmt"

func main() {
	var names = [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	var car [5]string

	car[0] = "Hyundai"
	car[1] = "Toyota"
	car[2] = "BYD"
	car[3] = "Honda"
	car[4] = "Tesla"
	fmt.Println(car)

	var number = [...]int{
		11,
		24,
		21,
		1,
		92,
		24,
		894,
	}

	fmt.Println("Panjang array", len(number), cap(number), number)

}
