package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan sebuah bilangan bulat positif: ")
	fmt.Scan(&n)

	if n < 2 {
		fmt.Println("bukan prima")
		return
	}

	prime := true
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			prime = false
			break
		}
	}

	if prime {
		fmt.Println("prima")
	} else {
		fmt.Println("bukan prima")
	}
}