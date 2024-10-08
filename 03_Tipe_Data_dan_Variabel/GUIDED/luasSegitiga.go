package main

import "fmt"

func main(){
	// mendeklarasikan variabel "alas" "tinggi" dan "hasil" dengan tipe data "float64"
	var alas float64
	var tinggi float64
	var hasil float64

	// menampilkan pesan yang meminta pengguna menginputkan alas segitiga
	fmt.Print("masukkan alas segitiga :")
	fmt.Scan(&alas)

	// menampilkan pesan yang meminta pengguna menginputkan tinggi segitiga
	fmt.Print("masukkan tinggi segitiga :")
	fmt.Scan(&tinggi)

	// rumus luas segitiga
	hasil = 0.5 * alas * tinggi

	// menampilkan luas segitiga
	fmt.Printf("luas segitiga adalah: %.2f", hasil)
}