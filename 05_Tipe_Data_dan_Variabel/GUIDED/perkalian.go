package main

import "fmt"

func main(){
	var i, a1, a2, hasil int

	// meminta pengguna untuk menginputkan dua angka
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scan(&a1)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scan(&a2)
	hasil = 0

	// melakukan perkalian tanpa operator "*"
	for i = 1; i <= a2; i+=1{
		hasil = hasil + a1
	}
	// menampilkan hasil perkalian 
	fmt.Println(hasil)
}