package main

import (
	"fmt"
)

func logging() {
	fmt.Println("End of program...")
}

func runApp() (name string) {
	defer logging()
	name = "Hary"
	fmt.Println("Starting application...")
	return name
}

func main() {
	runApp()
}
