// PSEUDOCODE
// 
// program PolaBilanganA
// kamus
//     x : integer
//     i, j : integer
// 
// algoritma
//     input(x)
// 
//     // Loop baris
//     for i ← 1 to x do
//         // Loop kolom
//         for j ← 1 to x do
//             output(j, " ")
//         endfor
//         output newline
//     endfor
// endprogram



// CODING
package main

import "fmt"

func main() {
    var x int
    fmt.Print("Masukkan bilangan bulat positif x: ")
    fmt.Scanln(&x)

    // Loop untuk baris
    for i := 1; i <= x; i++ {
        // Loop untuk kolom
        for j := 1; j <= x; j++ {
            fmt.Print(j, " ")
        }
        fmt.Println() // Pindah ke baris berikutnya
    }
}
