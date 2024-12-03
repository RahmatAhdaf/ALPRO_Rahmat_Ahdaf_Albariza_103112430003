package main

import (
	"fmt"
)

func main() {
	correctPassword := "12345"
	attempts := 0

	for attempts < 3 { // While-loop simulasi
		var input string
		fmt.Print("Masukkan password: ")
		fmt.Scanln(&input)

		if input == correctPassword {
			fmt.Println("Login berhasil")
			return
		}

		fmt.Println("Password salah")
		attempts++
	}

	fmt.Println("Login ditolak")
}