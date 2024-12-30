// PSEUDOCODE
// 
// program IsiTangki
// kamus
//     T : integer {kapasitas tangki}
//     V : integer {volume ember}
//     totalVolume : integer
//     penuh : boolean
// 
// algoritma
//     input(T) {kapasitas tangki}
//     totalVolume ← 0
//     penuh ← false
// 
//     while penuh = false do
//         input(V) {volume ember}
//         totalVolume ← totalVolume + V
//         if totalVolume >= T then
//             penuh ← true
//         else
//             penuh ← false
//         endif
//         output(penuh)
//     endwhile
// endprogram


// CODING
package main

import "fmt"


func main() {
	var T, totalVolume, volume int
	var penuh bool

	// Input kapasitas tangki
	fmt.Print("Masukkan kapasitas tangki (T): ")
	fmt.Scan(&T)

	totalVolume = 0
	penuh = false

	for !penuh {
		fmt.Print("Masukkan volume ember: ")
		fmt.Scan(&volume)
		totalVolume += volume
		if totalVolume >= T {
			penuh = true
		} else {
			penuh = false
		}
		fmt.Println(penuh)
	}
}
