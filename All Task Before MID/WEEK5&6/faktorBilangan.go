package main

import "fmt"


func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i, "true")
		} else {
			fmt.Println(i, "false")
		}
	}
}
