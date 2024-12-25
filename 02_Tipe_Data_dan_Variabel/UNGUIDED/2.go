package main

import "fmt"

func tampilkanBiodata(nama, nim, kelas string) {
	fmt.Printf("Perkenalkan saya adalah %s, salah satu\n", nama)
 	fmt.Printf("mahasiswa Prodi S1-IF dari kelas %s\n", kelas)
  	fmt.Printf("dengan NIM %s.\n", nim)
}

func main() {
	var nama, nim, kelas string

  	fmt.Print("Masukkan nama: ")
  	fmt.Scanln(&nama)

  	fmt.Print("Masukkan NIM: ")
  	fmt.Scanln(&nim)

  	fmt.Print("Masukkan kelas: ")
  	fmt.Scanln(&kelas)

  	tampilkanBiodata(nama, nim, kelas)
}