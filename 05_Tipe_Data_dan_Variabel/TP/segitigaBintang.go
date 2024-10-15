package main

import "fmt"

func main() {
    var n int

    // Meminta input jumlah baris
    fmt.Print("Masukkan jumlah baris: ")
    fmt.Scan(&n)

    // Loop untuk mencetak segitiga bintang
    for i := 1; i <= n; i++ {
        // Cetak bintang sesuai jumlah baris
        for j := 1; j <= i; j++ {
            fmt.Print("*")
        }
        // Pindah ke baris baru
        fmt.Println()
    }
}
