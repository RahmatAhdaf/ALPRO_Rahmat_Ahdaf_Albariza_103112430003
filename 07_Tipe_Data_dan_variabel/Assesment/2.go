package main

import "fmt"

func main(){
	var n int
	fmt.Print("Masukan nilai n: ")
	fmt.Scanln(&n)

	fmt.Print("Hasil kuadrat dari 1 sampai ", n ," : ")
	for i := 1; i <= n; i++{
		fmt.Print(i*i, " ")
	}
	fmt.Println()
}