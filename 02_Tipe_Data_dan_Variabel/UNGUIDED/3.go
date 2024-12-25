package main

import "fmt"

func main() {
	var jariJari float64

	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&jariJari)

	// Menghitung luas lingkaran (gunakan nilai π = 3.14159)
	luas := 3.14159 * jariJari * jariJari

	// Menampilkan hasil dengan 1 angka desimal
	fmt.Printf("Luas lingkaran adalah: %.1f\n", luas)
}
