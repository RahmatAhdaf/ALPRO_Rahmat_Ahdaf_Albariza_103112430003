package main

import "fmt"

func main(){
	var token string
	fmt.Print("Masukan token: ")
	fmt.Scan(&token)
	for token != "12345abcd" {
		fmt.Print("Masukan token: ")
		fmt.Scan(&token)
	}
	fmt.Println("Selamat Anda Berhasil Login")
}