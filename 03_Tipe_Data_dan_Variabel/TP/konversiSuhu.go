package main

import (
	"fmt"
)

// Fungsi untuk mengkonversi suhu dari Fahrenheit ke Kelvin
func fahrenheitToKelvin(fahrenheit float64) float64 {
	kelvin := (fahrenheit-32)*5/9 + 273.15
	return kelvin
}

func main() {
	var fahrenheit float64

	// input suhu dalam Fahrenheit
	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	// Mengonversi suhu ke Kelvin
	kelvin := fahrenheitToKelvin(fahrenheit)

	// Menampilkan hasil konversi
	fmt.Printf("Suhu dalam Kelvin: %.2f\n", kelvin)
}
