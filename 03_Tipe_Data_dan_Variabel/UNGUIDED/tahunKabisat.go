package main

import (
	"fmt"
)

func main() {
	// mendeklarasikan variabel "tahun" dengan tipe data "int"
	var tahun int

	// manampilkan pesan yang menyuruh pengguna menginputkan tahun
	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&tahun)

	// Logika untuk memeriksa apakah tahun kabisat
	if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0) {
		fmt.Println("kabisat : true") // Jika tahun kabisat
	} else {
		fmt.Println("kabisat : false") // Jika bukan tahun kabisat
	}
}
