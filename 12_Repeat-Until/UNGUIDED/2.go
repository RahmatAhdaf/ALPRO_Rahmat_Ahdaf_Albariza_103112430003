package main

import "fmt"

func main() {
	// Input: Bilangan desimal
	var number float64
	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&number)

	// Pembulatan ke atas (ceiling)
	roundedUp := float64(int(number))
	if number > roundedUp {
		roundedUp++
	}

	// Repeat-until style: menambahkan 0.1 setiap iterasi
	// Mulai dari nilai input yang sudah dibulatkan
	for number < roundedUp {
		// Menampilkan nilai setelah penambahan
		number += 0.1
		fmt.Printf("%.1f\n", number)
	}
}