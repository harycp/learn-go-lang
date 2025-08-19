package main

import "fmt"

func main() {
	var name = []string{
		"Hary",
		"Capri",
		"Sukro",
		"Makro",
		"Surtono",
		"Tartono",
		"Hendrawan",
		"Hensubroto",
		"Nabuoloto",
	}

	slice1 := name[3:9]
	slice2 := name[5:]
	var slice3 []string = name[:4]
	var slice4 []string = name[:]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	// Function Slice
	panjang := len(slice1)
	kapasitas := cap(slice1)

	fmt.Println(panjang)
	fmt.Println(kapasitas)

	var days = [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	daysSlice1 := days[0:7]
	daysSlice1[1] = "Selasa Baru"
	daysSlice1[2] = "Rabu Baru"
	fmt.Println(daysSlice1)

	daysSlice2 := append(daysSlice1, "Hari Baru") // Akan langsung membuat array baru jika kapasitas dari awal yang digunakan sudah maks
	fmt.Println(daysSlice2)
	fmt.Println(days)

	var newSlice []string = make([]string, 9, 11)
	for i := 0; i < 9; i++ {
		newSlice[i] = "Hari"
	}
	appendSlice := append(newSlice, "Hari")
	fmt.Println(newSlice)
	fmt.Println(appendSlice)
	newSlice = append(newSlice, "Hari")
	fmt.Println(newSlice)

}
