package main

import "fmt"

func main() {
	input := "tambah"
	var a, b, result int

	a = 2
	b = 9
	if input == "tambah" {
		result = a + b
	} else if input == "kurang" {
		result = a - b
	} else {
		result = a * b
	}

	fmt.Println("Input : ", input, "Hasil : ", result)
}
