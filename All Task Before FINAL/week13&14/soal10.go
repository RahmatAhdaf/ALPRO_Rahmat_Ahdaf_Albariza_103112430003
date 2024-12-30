// PSEUDOCODE
// 
// program PolaDiagonal
// kamus
//     x : integer {bilangan bulat positif untuk ukuran pola}
//     i, j : integer {indeks baris dan kolom}
//
// algoritma
//     input(x) {ukuran pola}
//     for i ← 0 to x-1 do
//         for j ← 0 to x-1 do
//             if i = j then
//                 output(j+1, " ")
//             else if i + j = x-1 then
//                 output(i+1, " ")
//             else
//                 output("  ")
//             endif
//         endfor
//         output(newline)
//     endfor
// endprogram



// CODING
package main

import "fmt"

func main(){
	var x int
	fmt.Scan(&x)

	for i:= 0; i < x; i++{
		for j := 0; j < x; j++{
			if i == j{
				fmt.Print(j+1, "  ")
			}else if i+j == x-1{
				fmt.Print(i+1, "  ")
			}else{
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
