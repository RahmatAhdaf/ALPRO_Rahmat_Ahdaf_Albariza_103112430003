package main

import "fmt"

func main() {

	// alun-alun purwokerto memiliki panjang sisi 27 meter
	sisi := 27.0

	// menghitung keliling persegi
	keliling := 4 * sisi

	// menghitung luas persegi
	luas := sisi * sisi

	// menampilkan output
	fmt.Printf("Panjang sisi alun-alun: %.2f meter\n", sisi)
	fmt.Printf("Keliling alun-alun: %.2f meter\n", keliling)
	fmt.Printf("Luas alun-alun: %.2f meter persegi\n", luas)
}