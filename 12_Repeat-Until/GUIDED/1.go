package main

import "fmt"

func main(){
	var kata string
	var perulangan int
	fmt.Print("Masukan kata dan jumlah perulangan: ")
	fmt.Scan(&kata, &perulangan)
	counter := 0
	for done := false; !done; {
		fmt.Println(kata)
		counter++
		done = (counter >= perulangan)
	}
}