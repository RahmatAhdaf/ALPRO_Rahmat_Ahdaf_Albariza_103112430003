package main

import (
	"fmt"
)

func main() {
	var n int
	// Meminta pengguna menginputkan jumlah kerucut
	fmt.Print("Masukan jumlah kerucut: ")
	// Membaca jumlah kerucut (n) dari input
	fmt.Scanln(&n)

	var r, t float64
	// Membuat slice untuk menyimpan volume setiap kerucut
	volume := make([]float64, n)

	for i := 0; i < n; i++ {
		// Meminta pengguna menginputkan jari-jari alas kerucut
		fmt.Print("Masukan jari-jari: ")
		// Membaca jari-jari (r) dari input
		fmt.Scanln(&r)
		// Meminta pengguna menginputkan tinggi kerucut
		fmt.Print("Masukan Tinggi: ")
		// Membaca tinggi (t) dari input
		fmt.Scanln(&t) // 
		volume[i] = (1.0/3.0)*3.14159*r*r*t // Menghitung volume dengan rumus (1/3) * π * r^2 * h
	}

	for _, v := range volume {
		// menampilkan hasil
		fmt.Println(v)
	}
}