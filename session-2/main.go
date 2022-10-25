package main

import "fmt"

func main() {
	// cara deklarasi pertama
	var firstName string = "john"

	// cara deklarasi kedua
	var lastName string
	lastName = "wick"

	// short declaration
	middleName := "robert"

	// println bisa pakai ",", bisa pakai "+" untuk concat string nya
	fmt.Println("first name:", firstName)
	fmt.Println("middle name: " + middleName)
	fmt.Println("last name: " + lastName)

	// understanding printf
	// %s untuk cetak string, %d untuk cetak integer, dll
	age := 10
	fmt.Printf("contoh string: %s \ncontoh integer: %d\n", middleName, age)

	// multi variable one line
	// kalau datanya ga mau diambil, bisa pakai "_" untuk dibuang dari memory
	first, second, _ := 1, 2, 3
	fmt.Printf("first: %d\nsecond: %d\n", first, second)
}
