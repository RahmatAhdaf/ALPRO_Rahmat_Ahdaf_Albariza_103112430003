package main

import "fmt"

func main() {
    var totalBelanja, diskon int

    // Input total belanja
    fmt.Print("Masukkan total belanja: ")
    fmt.Scan(&totalBelanja)

    // Input besar diskon
    fmt.Print("Masukkan besar diskon (dalam persen): ")
    fmt.Scan(&diskon)

    // Menghitung total setelah diskon
    totalAkhir := float64(totalBelanja) - (float64(totalBelanja) * float64(diskon) / 100)

    // Output total belanja akhir
    fmt.Printf("Total belanja setelah diskon: %.2f\n", totalAkhir)
}
