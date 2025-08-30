package main

import "fmt"

func desire() {
	fmt.Println("Choose a command")
}

func endOfProgram() {
	fmt.Println("End of program...")
	message := recover()
	fmt.Println("There's a message:", message)
}

func running(error bool) {
	defer endOfProgram()
	if error {
		panic("ERROR NIH BANG")
	}
	desire()
}

func main() {
	running(true)
}
