package main

import (
	"fmt"
)

func main() {
	// mendeklarasikan variabel "BMI" dan "TB" dengan tipe data "float64"
	var BMI, TB float64

	// Input BMI dan tinggi badan dalam meter
	fmt.Print("Masukkan nilai BMI: ")
	fmt.Scan(&BMI)
	fmt.Print("Masukkan tinggi badan dalam meter: ")
	fmt.Scan(&TB)

	// Menghitung berat badan
	BB := BMI * (TB * TB)

	// Menampilkan hasil berat badan
	fmt.Printf("Berat badan Anda adalah: %.2f kilogram\n", BB)
}
