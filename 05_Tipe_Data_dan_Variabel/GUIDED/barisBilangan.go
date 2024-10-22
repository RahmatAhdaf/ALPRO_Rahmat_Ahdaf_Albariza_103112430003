package main

import "fmt"

func main(){
	var a, b int

	// meminta pengguna untuk menginputkan nilai a dan b / bilangan awal dan akhir
	fmt.Print("Masukkan nilai a:")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b:")
	fmt.Scan((&b))
	for i := a; i <= b; i++ {
		// menampilkan hasil
		fmt.Print(i, )
	}
}