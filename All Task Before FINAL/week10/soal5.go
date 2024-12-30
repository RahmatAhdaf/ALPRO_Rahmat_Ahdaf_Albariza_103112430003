// PSEUDOCODE
// 
// program PromoMiniMarket
// kamus
//     totalBelanja : integer
//     buatKartu : boolean
//     diskon, cashback, totalBayar : integer
// 
// algoritma
//     input(totalBelanja, buatKartu)
//     if buatKartu then
//         kartu ← true
//     else
//         kartu ← false
//     endif
// 
//     if totalBelanja >= 100000 then
//         diskon ← totalBelanja * 10 / 100
//     else
//         diskon ← 0
//     endif
// 
//     if totalBelanja >= 200000 and kartu = true then
//         cashback ← 75000
//     else
//         cashback ← 0
//     endif
// 
//     totalBayar ← totalBelanja - diskon - cashback
//     output(kartu, diskon > 0, cashback > 0, totalBayar)
// endprogram


// CODING
package main

import "fmt"

func main() {
	// Input total belanja dan status pembuatan kartu
	var totalBelanja int
	var buatKartu bool

	fmt.Print("Masukkan total belanja: ")
	fmt.Scan(&totalBelanja)
	fmt.Print("Apakah ingin membuat kartu? (true/false): ")
	fmt.Scan(&buatKartu)

	// Tentukan status kartu
	kartu := buatKartu

	// Hitung diskon
	var diskon int
	if totalBelanja >= 100000 {
		diskon = totalBelanja * 10 / 100
	} else {
		diskon = 0
	}

	// Hitung cashback (hanya jika membuat kartu)
	var cashback int
	if totalBelanja >= 200000 && kartu {
		cashback = 75000
	} else {
		cashback = 0
	}

	// Hitung total bayar
	totalBayar := totalBelanja - diskon - cashback

	// Output hasil
	fmt.Printf("Kartu? %v\n", kartu)
	fmt.Printf("Diskon? %v\n", diskon > 0)
	fmt.Printf("Cashback? %v\n", cashback > 0)
	fmt.Printf("Total Bayar: Rp. %d\n", totalBayar)
}
