package main

import (
	"fmt"
	"math"
)

func main() {
	var kantong1, kantong2 float64

	for {
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantong1, &kantong2)

		if kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Berat tidak boleh negatif.")
			break
		}

		if kantong1+kantong2 > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		selisih := math.Abs(kantong1 - kantong2)
		fmt.Printf("Sepeda motor Pak Andi akan oleng: %t\n", selisih >= 9)
	}
}