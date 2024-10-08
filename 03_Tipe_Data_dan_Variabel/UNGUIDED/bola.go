package main

import (
	"fmt"
)

func main() {
	// mendeklarasikan variabel "r" dengan tipe data "float64"
	var r float64

	// manampilkan pesan yang menyuruh pengguna menginputkan jari-jari bola
	fmt.Print("masukkan jari-jari bola = ")
	fmt.Scan(&r)

	// Rumus untuk menghitung volume bola dan luas kulit bola
	volume :=  4 * 3.1415926535 * r * r * r / 3
	luas := 4 * 3.1415926535 * r * r

	// Menampilkan hasil
	fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}
