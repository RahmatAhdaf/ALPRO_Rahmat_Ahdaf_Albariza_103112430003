package main

import "fmt"

func main(){
	var dinar, dirham, fals, qirat int

	fmt.Print("Masukan satuan qirat: ")
	fmt.Scanln(&qirat)

	// konversi
	dinar = qirat /600
	dirham = qirat %600 /60
	fals = qirat %600 %60 /6
	qirat = qirat %600 %60 %6

	fmt.Printf("Satuan dinar: %d\n", dinar)
	fmt.Printf("Satuan dirham: %d\n", dirham)
	fmt.Printf("Satuan fals: %d\n", fals)
	fmt.Printf("Satuan qirat: %d", qirat)
}