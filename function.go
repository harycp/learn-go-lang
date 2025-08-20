package main

import "fmt"

// Function
func sayHello() {
	fmt.Println("Hello World")
}

// Function Parameter
func multiply(variable1 int, variable2 int) {
	result := variable1 * variable2
	fmt.Println(result)
}

// Function Return
func divide(variable1 int, variable2 int) int {
	result := variable1 / variable2
	return result
}

// Function multiple return
func oddEvenNumber(variable1 int, variable2 int) ([]int, []int) {
	var oddNumber, evenNumber []int

	if variable1%2 == 0 {
		evenNumber = append(evenNumber, variable1)
	} else {
		oddNumber = append(oddNumber, variable2)
	}

	if variable2%2 == 0 {
		evenNumber = append(evenNumber, variable2)
	} else {
		oddNumber = append(oddNumber, variable1)
	}

	return oddNumber, evenNumber
}

// Named return value

func main() {
	sayHello()
	multiply(4, 2)
	result := divide(4, 2)
	fmt.Println(result)
	odd, even := oddEvenNumber(4, 2)
	fmt.Println("Ganjil", odd)
	fmt.Println("Genap", even)

	//ignore return value
	odd1, _ := oddEvenNumber(9, 2)
	fmt.Println(odd1)
}
