package main

import "fmt"

func isSorted(num int) bool {
	// Mengambil digit ratusan, puluhan, dan satuan dari bilangan tiga digit.
	hundreds := num / 100         // Digit ratusan
	tens := (num / 10) % 10       // Digit puluhan
	units := num % 10             // Digit satuan

	// Memeriksa apakah digit terurut membesar atau mengecil.
	if hundreds <= tens && tens <= units {
		return true // Terurut membesar
	} else if hundreds >= tens && tens >= units {
		return true // Terurut mengecil
	} else {
		return false // Tidak terurut
	}
}

func main() {
	// Masukan angka untuk dicek.
	numbers := []int{149, 555, 961, 183}

	// Mengecek setiap angka dan menampilkan hasilnya.
	for _, num := range numbers {
		result := isSorted(num)
		fmt.Printf("Masukan: %d, Terurut: %t\n", num, result)
	}
}