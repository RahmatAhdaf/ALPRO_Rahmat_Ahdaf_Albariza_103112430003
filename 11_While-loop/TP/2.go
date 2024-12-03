package main

import "fmt"

func main() {
	var total float64
	var harga float64
	var pilihan string

	for {
		fmt.Print("\nMasukkan harga barang (atau ketik 'selesai' untuk mengakhiri): ")
		fmt.Scanln(&pilihan)

		if pilihan == "selesai" {
			break
		}

		fmt.Sscan(pilihan, &harga)
		total += harga
	}

	fmt.Printf("\nTotal belanja: %.2f\n", total)
	fmt.Println("Transaksi selesai. Terima kasih!")
}
