package main

import "fmt"

func main(){
	var alas, tinggi  float64
	var n int

	fmt.Print("Masukkan n:")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		// meminta pengguna untuk menginputkan alas dan tinggi segitiga
		fmt.Print("Masukkan alas:")
		fmt.Scan(&alas)
		fmt.Print("Masukkan tinggi: ")
		fmt.Scan(&tinggi)

		// rumus luas segitiga
		luas := alas * tinggi / 2

		// menampilkan hasil
		fmt.Println("Luas: ", luas,"    ")
	}
}