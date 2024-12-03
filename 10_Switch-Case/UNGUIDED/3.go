package main

import "fmt"

func main() {
	var input int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&input)

	switch {
	case input%2 != 0:
		// Bilangan Ganjil
		hasil := input + (input + 1)
		fmt.Println("Kategori: Bilangan Ganjil")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", input, input+1, hasil)

	case input%2 == 0 && input%10 != 0:
		// Bilangan Genap
		hasil := input * (input + 1)
		fmt.Println("Kategori: Bilangan Genap")
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n", input, input+1, hasil)

	case input%5 == 0 && input%10 != 0:
		// Bilangan Kelipatan 5
		hasil := input * input
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		fmt.Printf("Hasil kuadrat dari %d = %d\n", input, hasil)

	case input%10 == 0:
		// Bilangan Kelipatan 10
		hasil := input / 10
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", input, hasil)

	default:
		fmt.Println("Bilangan tidak termasuk dalam kategori yang ditentukan.")
	}
}