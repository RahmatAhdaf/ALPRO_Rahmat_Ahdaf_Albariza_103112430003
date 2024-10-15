package main

import "fmt"


func main() {
	// mendeklarasikan variabel "jam" "menit" "detik" dan "masukkan" dengan tipe data "int"
	var jam, menit, detik, masukkan int
	// mengambil input dari pengguna dan menyimpannya ke variabel masukkan
	fmt.Scanln(&masukkan)
	// rumus konversi detik ke jam, menit, dan detik
	jam = masukkan/3600
	menit = (masukkan%3600)/60
	detik = masukkan %60
	// menampilkan output hasil konversi
	fmt.Println(jam, "jam", menit, "menit", detik, "detik")

}