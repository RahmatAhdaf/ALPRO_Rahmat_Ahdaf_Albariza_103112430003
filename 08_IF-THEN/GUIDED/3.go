package main

import "fmt"

func main(){
	var bilangan int
	var b bool
	fmt.Print("masukan bilangan: ")
	fmt.Scan(&bilangan)

	if bilangan%2 == 0 && bilangan<0 {
		b = true
	}
	fmt.Print(b)
}