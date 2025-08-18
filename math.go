package main

import "fmt"

func main() {
	a := 2
	b := 4
	c := 10
	d := 9
	e := 7
	f := 2

	result := (a + (b * c) - d) / e

	modulus := result % f
	fmt.Println(result)
	fmt.Println(modulus)
}
