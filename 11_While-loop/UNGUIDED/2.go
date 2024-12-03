package main

import "fmt"

func main() {
	var input int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scanln(&input)

	fmt.Println("Output digit (dari kanan ke kiri):")

	// While loop menggunakan for
	for input > 0 {
		fmt.Println(input % 10) // Cetak digit terakhir
		input /= 10            // Hapus digit terakhir
	}
}
