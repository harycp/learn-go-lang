package main

import "fmt"

func main() {
	a := "Hary"
	b := "Hary"
	c := 2
	d := 8

	equal := a == b
	notEqual := a != b
	greater := c > d
	less := c < d
	greaterEqual := c >= d
	lessEqual := c <= d

	fmt.Println(equal)
	fmt.Println(notEqual)
	fmt.Println(greater)
	fmt.Println(less)
	fmt.Println(greaterEqual)
	fmt.Println(lessEqual)

}
