package main

import "fmt"

func main(){
	var dinar, dirham, fals, qirat float64

	fmt.Print("Masukan satuan qirat: ")
	fmt.Scanln(&qirat)

	// konversi
	dinar = qirat/600
	dirham = qirat/60
	fals = qirat/6

	fmt.Printf("Satuan dinar: %.2f\n", dinar)
	fmt.Printf("Satuan dirham: %.2f\n", dirham)
	fmt.Printf("Satuan fals: %.2f", fals)
}