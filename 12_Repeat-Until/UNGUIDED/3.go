package main

import "fmt"

func main() {
	// Input: Target donasi
	var target int
	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)

	// Variabel untuk menghitung total donasi dan jumlah donatur
	var totalDonasi, jumlahDonatur int

	// Repeat-until style: program terus meminta input hingga target tercapai
	for {
		var donasi int
		fmt.Print("Masukkan donasi dari donatur: ")
		fmt.Scan(&donasi)

		// Menambahkan donasi ke total
		totalDonasi += donasi
		jumlahDonatur++

		// Menampilkan total donasi dan jumlah donatur setiap iterasi
		fmt.Printf("Total donasi: %d, Jumlah donatur: %d\n", totalDonasi, jumlahDonatur)

		// Kondisi berhenti (repeat-until): total donasi >= target
		if totalDonasi >= target {
			break
		}
	}

	// Output akhir setelah target tercapai
	fmt.Println("Target donasi tercapai!")
	fmt.Printf("Total donasi akhir: %d\n", totalDonasi)
	fmt.Printf("Jumlah donatur: %d\n", jumlahDonatur)
}