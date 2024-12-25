package main

import "fmt"

func main() {
	var fahrenheit float64

	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	// Menghitung suhu dalam Celsius
	celsius := (fahrenheit - 32) * 5 / 9

	// Menampilkan hasil
	fmt.Printf("Suhu dalam Celsius adalah: %.1f\n", celsius)
}
