package main

import "fmt"

func main() {
	var option string
	var a, b, result float64

	option = "bagi"
	a = 9
	b = 4
	switch option {
	case "tambah":
		result = a + b
	case "kurang":
		result = a - b
	case "bagi":
		result = a / b
	case "kali":
		result = a * b
	default:
		result = 0
	}

	fmt.Println(result)

}
