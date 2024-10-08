package main

import "fmt"

func main() {
	// mendeklarasikan variabel "rupiah" dan "usdollar" dengan tipe data "float64"
    var rupiah float64
	var usdollar float64
    
	//menampilkan pesan yang meminta pengguna menginputkan satuan rupiah
    fmt.Print("Masukkan satuan rupiah: ")
    fmt.Scan(&rupiah)

	// konversi rupiah ke USD
	usdollar = rupiah/15000

	// menampilkan hasil konversi
	fmt.Printf("Jumlah dalam USD: %.2f USD\n", usdollar)
}
