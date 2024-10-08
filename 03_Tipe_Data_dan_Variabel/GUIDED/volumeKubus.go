package main

import "fmt"

func main(){
	// mendeklarasikan variabel "sisi" dengan tipe data "float64" 
	var sisi float64
	// mendeklarasikan variabel "hasil" dengan tipe data "float64"
	var hasil float64

	// Menampilkan pesan yang meminta pengguna menginputkan panjang kubus
	fmt.Print("masukkan panjang sisi kubus:")
	fmt.Scan(&sisi)

	// rumus volume kubus
	hasil = sisi * sisi * sisi

	// menampilkan volume kubus
	fmt.Printf("volume kubusnya adalah: %.2f meter kubik\n", hasil)
}