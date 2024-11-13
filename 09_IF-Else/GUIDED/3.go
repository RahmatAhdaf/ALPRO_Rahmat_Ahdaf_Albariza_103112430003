package main

import (
	"fmt"
)

func main() {
	var bilangan int

	fmt.Print("Masukkan bilangan 4 digit (1000 - 9999): ")
	fmt.Scan(&bilangan)

	// Memeriksa apakah bilangan memenuhi syarat 4 digit
	if bilangan < 1000 || bilangan > 9999 {
		fmt.Print("Bilangan harus 4 digit dan berada antara 1000 hingga 9999.")
		return
	}

	digit1 := bilangan / 1000            // ribuan
	digit2 := (bilangan / 100) % 10      // ratusan
	digit3 := (bilangan / 10) % 10       // puluhan
	digit4 := bilangan % 10              // satuan

	if digit1 < digit2 && digit2 < digit3 && digit3 < digit4 {
		fmt.Println("Digit terurut membesar")
	} else if digit1 > digit2 && digit2 > digit3 && digit3 > digit4 {
		fmt.Println("Digit terurut mengecil")
	} else {
		fmt.Println("Digit tidak terurut")
	}
}