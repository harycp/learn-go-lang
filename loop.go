package main

import "fmt"

func main() {
	var count, total int
	count = 1

	for count <= 10 {
		fmt.Println(count)
		count++
	}

	for total = 0; total < count; total++ {
		if total%2 == 0 {
			fmt.Println("Genap : ", total)
		}
		fmt.Println("Ganjil : ", total)
	}

	// Slice
	names := []string{"Alice", "Bob", "Charlie", "Ucup"}
	var i int
	for i = 0; i < len(names); i++ {
		fmt.Println("Index ke-", i, "=", names[i])
	}
	for index, name := range names {
		fmt.Println("Index ke-", index, "=", name)
	}
	for _, name := range names {
		fmt.Println("names = ", name)
	}

}
