package main

import "fmt"

func main() {
	var jumlahBarang int
	// Meminta pengguna menginputkan jumlah barang yang telah dibeli
	fmt.Print("Masukkan jumlah barang yang dibeli: ")
	fmt.Scanln(&jumlahBarang)

	var jumlahPoin int = 0

	// Menghitung poin normal
	jumlahPoin = jumlahBarang * 10

	// Menghitung poin tambahan
	if jumlahBarang > 5 {
		jumlahPoin += (jumlahBarang - 5) * 5
	}
	// Menampilkan hasil jumlah poin yang didapatkan
	fmt.Printf("Total poin yang didapatkan: %d poin\n", jumlahPoin)
}