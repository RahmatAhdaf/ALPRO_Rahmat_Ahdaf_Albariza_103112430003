package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan dua bilangan (x y): ")
	fmt.Scanln(&x, &y)

	result := 0

	// Menghitung hasil integer division menggunakan pengurangan berulang
	for x >= y {
		x -= y
		result++
	}

	fmt.Println("Hasil:", result)
}
