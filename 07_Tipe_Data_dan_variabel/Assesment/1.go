package main

import "fmt"

func main(){
	var x, y, i int
	fmt.Print("Masukan angka pertama: ")
	fmt.Scanln(&x)
	fmt.Print("Masukan angka kedua: ")
	fmt.Scanln(&y)

	sum := 0
	for i = x; i <= y; i++{
		sum += i
	}
	fmt.Printf("Hasil penjumlahannya adalah: %d ", sum )
}