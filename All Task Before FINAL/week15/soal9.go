// PSEUDOCODE
// 
// program Novel
// variabel
//     P, R : integer
//     hari : integer
// algoritma
//     input(P, R)
//     if R >= P then
//         hari = 1
//     else 
//         hari = P div R
//     	if P mod R != 0 then
//         	hari = hari + 1
//     endif
//     output(hari, "hari")
// endprogram


// CODING
package main

import "fmt"

func main(){
    var P, R int
    var hari int

    fmt.Scan(&P, &R)
    if R >= P{
		hari = 1
    } else {
        hari = P / R
        if P%R != 0{
            hari = hari + 1
        }
    }

    fmt.Printf("%d hari\n", hari)
}