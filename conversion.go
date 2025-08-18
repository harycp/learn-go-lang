package main

import "fmt"

func main() {
	var a int32 = 90
	var b = int64(a)

	fmt.Println(b)

	var name string = "Hary Capri"
	var pick = name[2]

	var conversion = string(pick)
	fmt.Println(conversion)
}
