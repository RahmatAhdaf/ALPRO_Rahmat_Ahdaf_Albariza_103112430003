package main

import (
	"fmt"
)

func main() {
	var orang int
	fmt.Print("Masukkan jumlah orang yang akan touring: ")
	fmt.Scan(&orang)

	var motor int

	// Menggunakan IF untuk menentukan jumlah motor yang diperlukan
	if orang%2 == 0 {
		// Jika jumlah orang genap, bagi dengan 2
		motor = orang / 2
	} else {
		// Jika jumlah orang ganjil, bagi dengan 2 dan tambahkan 1
		motor = (orang / 2) + 1
	}

	fmt.Printf("Jumlah motor yang diperlukan: %d\n", motor)
}