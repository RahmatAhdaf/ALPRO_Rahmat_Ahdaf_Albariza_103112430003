package main

import "fmt"

func main() {
	var bil1, pangkat int
	// Meminta input dari pengguna untuk bilangan yang akan dipangkat
	fmt.Print("Masukkan bilangan pertama: ")
	fmt.Scanln(&bil1)
	// Meminta input dari pengguna untuk pangkat
	fmt.Print("Masukkan bilangan kedua(pangkat): ")
	fmt.Scanln(&pangkat)

	// Inisialisasi hasil dengan 1
	hasil := 1
	// Looping sebanyak pangkat kali
	for i := 0; i < pangkat; i++ { 
		// Mengalikan hasil dengan basis pada setiap iterasi
		hasil *= bil1
	}
	// Menampilkan hasil pemangkatan
	fmt.Println("Hasil pemangkatan:", hasil)
}