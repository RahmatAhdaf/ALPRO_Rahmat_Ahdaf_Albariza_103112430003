package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik dalam koordinat 2D
func jarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func main() {
	// mendeklarasikan variabel "xA" "yA" "xB" "yB" "xC" dan "yC" dengan tipe data "float64"
	var xA, yA, xB, yB, xC, yC float64

	// Meminta input koordinat titik A, B, dan C
	fmt.Print("Masukkan koordinat titik A (x y): ")
	fmt.Scan(&xA, &yA)
	fmt.Print("Masukkan koordinat titik B (x y): ")
	fmt.Scan(&xB, &yB)
	fmt.Print("Masukkan koordinat titik C (x y): ")
	fmt.Scan(&xC, &yC)

	// Menghitung panjang sisi-sisi segitiga
	AB := jarak(xA, yA, xB, yB)
	BC := jarak(xB, yB, xC, yC)
	CA := jarak(xC, yC, xA, yA)

	// Menentukan sisi terpanjang
	terpanjang := math.Max(AB, math.Max(BC, CA))

	// Menampilkan hasil sisi terpanjang dengan dua angka di belakang koma
	fmt.Printf("Sisi terpanjang dari segitiga adalah: %.2f\n", terpanjang)
}
