package main

import "fmt"

func main() {
	var name string = "Hary"
	var age string = "Dua Puluh Satu"

	fmt.Println("Nama saya ", name)
	fmt.Println("Umur saya ", age)

	var panjang int = len(name)
	fmt.Println("Panjang nama saya ", panjang)

	var kedua int = int(name[1]) // dapet byte nya bukan stringnya
	fmt.Println("Huruf kedua nama saya ", kedua)

}
