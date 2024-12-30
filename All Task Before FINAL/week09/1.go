package main

import "fmt"

func main(){
	fmt.Println("===pertemuan 9===")

	input := 11

	if input > 10{
		fmt.Println(input, "lebih besar dari 10")
	} else if input >= 10{
		fmt.Println(input, "sama dengan 10")
	} else{
		fmt.Println(input, "lebih kecil dari 10")
	}
}