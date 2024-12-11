package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	// Input: tiga bilangan bulat
	fmt.Print("Masukkan tiga bilangan bulat (pisahkan dengan spasi):")
	fmt.Scan(&a, &b, &c)

	// Menentukan bilangan terbesar dan terkecil
	max := a
	min := a

	if b > max {
		max = b
	}
	if c > max {
		max = c
	}

	if b < min {
		min = b
	}
	if c < min {
		min = c
	}

	// Output: bilangan terbesar dan terkecil
	fmt.Printf("Bilangan terbesar: %d\n", max)
	fmt.Printf("Bilangan terkecil: %d\n", min)
}
