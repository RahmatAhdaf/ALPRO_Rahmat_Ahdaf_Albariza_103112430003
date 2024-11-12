package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&bilangan)

	if bilangan < 0 {
		bilangan = -bilangan
	}

	fmt.Print("Nilai absolutnya adalah: ", bilangan)
}
