// PSEUDOCODE
//
// program Konsekutif
// kamus
//     X : integer
//     isKonsekutif : boolean
//     prevDigit, currentDigit : integer
// algoritma
//     input(X)                                    
//     isKonsekutif ← true                         
//     prevDigit ← -1                              

//     selama X > 0 lakukan                       
//         currentDigit ← X % 10                   
//         jika prevDigit ≠ -1 dan |currentDigit - prevDigit| ≠ 1 maka
//             isKonsekutif ← false                
//             keluar                              
//         endif
//         prevDigit ← currentDigit                
//         X ← X div 10                            
//     akhir selama

//     output(isKonsekutif)                        
// endprogram



// CODING
package main

import "fmt"

func main() {
	// Input bilangan sebagai integer
	var X int
	fmt.Print("Masukkan bilangan positif: ")
	fmt.Scan(&X)

	// Variabel untuk menentukan apakah bilangan konsekutif
	isKonsekutif := true
	prevDigit := -1 // Nilai awal untuk digit sebelumnya

	// Periksa setiap digit dari kanan ke kiri
	for X > 0 {
		// Ambil digit paling kanan
		currentDigit := X % 10

		// Jika bukan digit pertama, cek selisihnya dengan digit sebelumnya
		if prevDigit != -1 && (currentDigit-prevDigit != 1 && prevDigit-currentDigit != 1) {
			isKonsekutif = false
			break
		}

		// Update digit sebelumnya dan buang digit paling kanan
		prevDigit = currentDigit
		X /= 10
	}

	// Output hasil
	fmt.Println(isKonsekutif)
}
