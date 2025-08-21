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
func completeData() (name string, age int, address string) {
	name = "Hary"
	age = 21
	address = "Limo simpang lima tiga ratus"

	return name, age, address
}

// Variadic Function
func multipleAll(numbers ...int) int {
	total := 1
	for _, number := range numbers {
		total *= number
	}
	return total
}

// Function value
func sum(numbers ...int) int {
	total := 0
	for _, numbers := range numbers {
		total += numbers
	}

	return total
}

// Function as Parameters

// Type function declaration
type Filter func(string) string

func filterWords(words string, filter Filter) string {
	wordsFilter := filter(words)
	return wordsFilter
}

func filter(words string) string {
	if words == "Anjing" || words == "Babi" || words == "Setan" {
		return "...."
	}
	return words
}

// Anonymous Function
type BlockedList func(string) bool

func welcomeHello(name string, blocked BlockedList) string {
	if blocked(name) {
		return "You are Blocked " + name
	} else {
		return "Welcome on Board " + name
	}
}

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

	// named return value
	name, age, address := completeData()
	fmt.Println(name, age, address)

	// variadic function
	total := multipleAll(2, 4, 1, 5, 23, 28, 2, 9, 50)
	fmt.Println(total)

	// slices parametes
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(multipleAll(numbers...))

	// Function Value
	sumAll := sum
	fmt.Println(sumAll(2, 4, 6, 23, 5))

	// Function as parameters
	fmt.Println(filterWords("Hary Capri", filter))

	filterName := filterWords("Anjing", filter)
	fmt.Println(filterName)

	// Anonymous Function
	blocked := func(word string) bool {
		return word == "Anjing" || word == "Babi" || word == "Setan"
	}
	fmt.Println(welcomeHello("Hary Capri ", blocked))
	fmt.Println(welcomeHello("Anjing", blocked))
}
