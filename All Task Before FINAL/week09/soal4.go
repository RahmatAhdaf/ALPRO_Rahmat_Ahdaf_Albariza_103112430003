// PSEUDOCODE
// 
// program Kelipatan
// kamus
//     n : integer
// 
// algoritma
//     input(n)
//     if n mod 3 = 0 and n mod 5 = 0 then
//         output("Kelipatan 3 dan Kelipatan 5")
//     else if n mod 3 = 0 then
//         output("Kelipatan 3")
//     else if n mod 5 = 0 then
//         output("Kelipatan 5")
//     endif
// endprogram


// CODING
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan sebuah bilangan bulat positif: ")
	fmt.Scan(&n)

	if n%3 == 0 && n%5 == 0 {
		fmt.Println("Kelipatan 3 dan Kelipatan 5")
	} else if n%3 == 0 {
		fmt.Println("Kelipatan 3")
	} else if n%5 == 0 {
		fmt.Println("Kelipatan 5")
	}
}
