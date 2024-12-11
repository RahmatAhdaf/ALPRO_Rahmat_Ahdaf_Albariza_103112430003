package main

import (
	"fmt"
)

func main() {
	var n int

	// Input: satu bilangan bulat positif
	fmt.Print("Masukkan sebuah bilangan bulat positif: ")
	fmt.Scan(&n)

	// Validasi input
	if n <= 0 {
		fmt.Println("Harap masukkan bilangan bulat positif.")
		return
	}

	// Menentukan faktor-faktor bilangan
	fmt.Printf("Faktor-faktor dari %d adalah:\n", n)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
	}
}
