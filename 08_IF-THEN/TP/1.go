package main

import "fmt"

func main() {
	var nilaiUjian int

	// meminta pengguna menginputkan nilai ujian
	fmt.Print("Nilai Ujian: ")
	fmt.Scan(&nilaiUjian)

	// jika nilai tersebut lebih besar atau sama dengan 70 menampilkan pesan "Lulus"
	if nilaiUjian >= 70 {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Tidak Lulus")
	}
}
