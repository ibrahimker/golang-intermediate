package main

import "fmt"

func main() {
	fmt.Println(GetLuasKubus(6))
	fmt.Println(GetUser())
}

func GetLuasKubus(sisi int) int {
	if sisi == 8 {
		return 8 * 8 * 8
	}
	return sisi * sisi * sisi
}

func GetUser() []string {
	return []string{"andi", "budi"}
}
