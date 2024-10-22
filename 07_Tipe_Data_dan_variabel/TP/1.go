package main

import "fmt"

func main() {
	var n int
	// Meminta pengguna menginputkan nilai N
	fmt.Print("Masukkan nilai N: ")
	// Membaca input dari pengguna dan menyimpannya ke variabel n
	fmt.Scanln(&n)

	// Menampilkan teks sebelum mencetak hasil kuadrat
	fmt.Print("Hasil kuadrat dari 1 sampai ", n, ": ")
	// Looping dari 1 hingga N
	for i := 1; i <= n; i++ {
		// Mencetak kuadrat dari i, yaitu i * i, diikuti spasi
		fmt.Print(i*i, " ")
	}
	// Menambahkan baris baru di akhir output
	fmt.Println()
}