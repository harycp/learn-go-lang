package main

import "fmt"

func main() {

	// break
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println(i)
			break
		}
	}

	// continue
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
	}
}
