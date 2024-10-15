package main

import "fmt"

func main() {
    var n int

    // Input nilai n
    fmt.Print("Masukkan nilai n: ")
    fmt.Scan(&n)

    // Inisialisasi variabel untuk menyimpan jumlah
    total := 0

    // Menjumlahkan deret angka dari 1 hingga n
    for i := 1; i <= n; i++ {
        total += i
    }

    // Output hasil penjumlahan
    fmt.Printf("Jumlah deret dari 1 hingga %d adalah: %d\n", n, total)
}
