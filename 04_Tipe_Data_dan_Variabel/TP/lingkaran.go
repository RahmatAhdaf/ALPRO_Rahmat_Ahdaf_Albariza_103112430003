package main

import "fmt"

func main() {
	// mendeklarasikan variabel "r" dengan tipe data "float64"
    var r float64
	// mendeklarasikan konstanta "pi" adalah 3.145926535
    const pi = 3.1415926535

    // menampilkan pesan yang meminta pengguna memasukkan jari-jari lingkaran
    fmt.Print("Masukkan jari-jari lingkaran: ")
    fmt.Scanln(&r)

    // rumus menghitung luas lingkaran
    luas := pi * r * r

    // rumus menghitung keliling lingkaran
    keliling := 2 * pi * r

    // Menampilkan hasil
    fmt.Printf("Luas lingkaran: %.2f\n", luas)
    fmt.Printf("Keliling lingkaran: %.2f\n", keliling)
}