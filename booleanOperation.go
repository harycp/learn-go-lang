package main

import "fmt"

func main() {
	var option1, option2, option3, option4, option5, option6, option7, option8, option9 bool
	var benar bool = true
	var salah bool = false

	option1 = benar && benar
	option2 = benar && salah
	option3 = salah && benar
	option4 = salah && benar

	option5 = benar || benar
	option6 = benar || salah
	option7 = salah || benar
	option8 = salah || salah

	option9 = !(benar && benar)
	
	fmt.Println(option1)
	fmt.Println(option2)
	fmt.Println(option3)
	fmt.Println(option4)
	fmt.Println(option5)
	fmt.Println(option6)
	fmt.Println(option7)
	fmt.Println(option8)
	fmt.Println(option9)

}
