package main

import "fmt"

func main() {
	// mendeklarasikan variabel "fx" dengan tipe data "float64"
    var fx float64

    // menampilkan pesan yang menyuruh pengguna menginputkan nilai f(x)
    fmt.Print("Masukkan nilai f(x): ")
    fmt.Scan(&fx)

    // Menghitung dan menampilkan nilai x
    x := (2 / (fx - 5)) - 5
    fmt.Printf("Nilai x adalah: %.2f\n", x)
    
}
