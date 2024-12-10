package main

import "fmt"

// Fungsi untuk menentukan apakah sebuah bilangan merupakan bilangan sempurna
func isPerfectNumber(num int) bool {
	if num <= 0 {
		return false
	}
	sum := 0
	// Menjumlahkan semua faktor bilangan kecuali bilangan itu sendiri
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	// Bilangan sempurna jika jumlah faktor sama dengan bilangan itu sendiri
	return sum == num
}

func main() {
	var input int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&input)

	// Mengevaluasi apakah bilangan tersebut sempurna
	if isPerfectNumber(input) {
		fmt.Printf("Ya (karena faktor dari %d adalah bilangan yang jumlahnya sama dengan %d)\n", input, input)
	} else {
		fmt.Printf("Tidak (bukan bilangan sempurna)\n")
	}
}