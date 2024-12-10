package main

import "fmt"

func main() {
	// Input: Bilangan bulat positif
	var number int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&number)

	// Cek apakah bilangan positif
	if number <= 0 {
		fmt.Println("Masukkan bilangan bulat positif.")
		return
	}

	// Menghitung banyaknya digit
	digitCount := 0
	for number > 0 {
		number /= 10 // Menghapus digit terakhir
		digitCount++ // Menambah jumlah digit
	}

	// Output: Banyaknya digit
	fmt.Printf("Banyaknya digit adalah %d\n", digitCount)
}