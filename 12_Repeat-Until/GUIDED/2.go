package main

import "fmt"

func main(){
	var num int
	var contLoop bool
	for contLoop = true; contLoop; {
		fmt.Print("Masukan bilangan bulat: ")
		fmt.Scan(&num)
		contLoop = num <= 0
	}
	fmt.Printf("%d adalah bilangan bulat positif\n", num)
}