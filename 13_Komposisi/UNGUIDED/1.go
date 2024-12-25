package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan sebuah bilangan bulat positif: ")
	fmt.Scan(&n)

	count := 0
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			count++
		}
	}

	fmt.Println("Terdapat", count, "bilangan ganjil")
}