package main

import "fmt"

func main() {
	var n int
	// Meminta pengguna untuk memasukkan bilangan bulat positif
	fmt.Print("Masukkan bilangan bulat positif: ")
	// Membaca input dari pengguna dan menyimpannya ke variabel 'n'
	fmt.Scanln(&n) 

	// Inisialisasi variabel 'sum' dengan nilai 0 untuk menyimpan hasil penjumlahan
	sum := 0 
	for i := 1; i <= n; i++ { 
		// Menambahkan nilai 'i' ke variabel 'sum' pada setiap iterasi loop
		sum += i 
	}
	// Menampilkan hasil penjumlahan
	fmt.Printf("Hasil penjumlahan dari 1 sampai %d adalah: %d\n", n, sum)
}