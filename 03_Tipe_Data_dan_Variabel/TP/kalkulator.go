package main

import (
	"fmt"
)

// Fungsi untuk penjumlahan
func tambah(a, b float64) float64 {
	return a + b
}

// Fungsi untuk pengurangan
func kurang(a, b float64) float64 {
	return a - b
}

// Fungsi untuk perkalian
func kali(a, b float64) float64 {
	return a * b
}

// Fungsi untuk pembagian
func bagi(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan.")
		return 0
	}
	return a / b
}

func main() {
	var angka1, angka2 float64
	var operasi string

	// Input dari pengguna
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&angka1)

	fmt.Print("Masukkan operasi (+, -, *, /): ")
	fmt.Scanln(&operasi)

	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&angka2)

	//untuk menentukan operasi
	switch operasi {
	case "+":
		fmt.Printf("Hasil: %.2f\n", tambah(angka1, angka2))
	case "-":
		fmt.Printf("Hasil: %.2f\n", kurang(angka1, angka2))
	case "*":
		fmt.Printf("Hasil: %.2f\n", kali(angka1, angka2))
	case "/":
		fmt.Printf("Hasil: %.2f\n", bagi(angka1, angka2))
	default:
		fmt.Println("Operasi tidak valid. Silakan gunakan +, -, *, atau /.")
	}
}
