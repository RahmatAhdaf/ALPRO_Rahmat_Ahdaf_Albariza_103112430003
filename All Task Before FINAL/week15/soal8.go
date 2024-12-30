// PSEUDOCODE
// 
// program sportClub
// kamus
//    S : integer  
//    minggu : integer  
//    hariKe : integer  
//
// algoritma
//     input(S)  
// 
//     if S = 0 then
//        output "Minggu ke-0"
//        output "Tidak memiliki jatah renang"
//     else
//         minggu ← S div 7  
//         hariKe ← S mod 7  
// 
//         if hariKe > 0 then
//             minggu ← minggu + 1  
//         endif
// 
//         output "Minggu ke-", minggu
//     endif
// endprogram



// CODING
package main

import "fmt"

func main() {
    var S int
    fmt.Print("Masukkan jatah berenang maksimum: ")
    fmt.Scanln(&S)

    if S == 0 {
        fmt.Println("Minggu ke-0")
        fmt.Println("Tidak memiliki jatah renang")
    } else {
        minggu := S / 7       // Hitung jumlah minggu penuh
        hariKe := S % 7       // Sisa hari setelah minggu penuh

        if hariKe > 0 {
            minggu++          // Tambah minggu jika ada sisa hari
        }

        fmt.Printf("Minggu ke-%d\n", minggu)
    }
}
