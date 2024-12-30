// PSEUDOCODE
//
// program JumlahGol
// kamus
//     a, b, c, d : integer
//     maxGol, minGol : integer
// 
// algoritma
//     input(a, b, c, d)
//     maxGol ← a
//     if b > maxGol then
//         maxGol ← b
//     endif
//     if c > maxGol then
//         maxGol ← c
//     endif
//     if d > maxGol then
//         maxGol ← d
//     endif
// 
//     minGol ← a
//     if b < minGol then
//         minGol ← b
//     endif
//     if c < minGol then
//         minGol ← c
//     endif
//     if d < minGol then
//         minGol ← d
//     endif
// 
//     output(maxGol, minGol)
// endprogram


// CODING
package main

import "fmt"

func main() {
	// Input jumlah gol dari empat tim
	var a, b, c, d int
	fmt.Println("Masukkan jumlah gol untuk empat tim (pisahkan dengan spasi):")
	fmt.Scan(&a, &b, &c, &d)

	// Menentukan jumlah gol maksimum dan minimum
	maxGol := a
	if b > maxGol {
		maxGol = b
	}
	if c > maxGol {
		maxGol = c
	}
	if d > maxGol {
		maxGol = d
	}

	minGol := a
	if b < minGol {
		minGol = b
	}
	if c < minGol {
		minGol = c
	}
	if d < minGol {
		minGol = d
	}

	// Output hasil
	fmt.Println("Jumlah gol terbanyak:", maxGol)
	fmt.Println("Jumlah gol tersedikit:", minGol)
}