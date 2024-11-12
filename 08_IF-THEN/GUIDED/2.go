package main

import "fmt"

func main(){
	var bilangan int
	fmt.Print("Masukan bilangan: ")
	fmt.Scan(&bilangan)

	if bilangan < 1 {
		fmt.Print("bukan positif")
	} else {
		fmt.Print("positif")
	}
}