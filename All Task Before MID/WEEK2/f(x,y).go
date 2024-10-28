package main

import (
	"fmt"
)

// Fungsi untuk menghitung f(x, y)
func f(x int, y int) float64 {
	return (1.0 / (3*float64(x)*float64(x))) + (10*float64(y)) + 7
}

func main() {
	// Contoh masukan
	var x, y int

	// Input dari pengguna
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	// Menghitung hasil dari f(x, y)
	hasil := f(x, y)

	// Menampilkan hasil
	fmt.Printf("Hasil dari f(%d, %d) = %f\n", x, y, hasil)
}
