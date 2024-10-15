package main

import "fmt"

func main() {
	// mendeklarasikan variabel "TB" "BB" dan "BMI" dengan tipe data "float64"
	var TB, BB, BMI float64

	// meminta pengguna menginputkan berat badan
	fmt.Print("masukkan BB:")
	fmt.Scan(&BB)
	// meminta pengguna menginputkan tinggi badan
	fmt.Print("masukkan TB:")
	fmt.Scan(&TB)
	// rumus menghitung body mass indeks adalah  hasil bagi dari berat badan dengan kuadrat dari tinggi badan.
	BMI = BB / (TB * TB)
	// menampilkan output dari body mass indeks
	fmt.Printf("%.2f", BMI)
}