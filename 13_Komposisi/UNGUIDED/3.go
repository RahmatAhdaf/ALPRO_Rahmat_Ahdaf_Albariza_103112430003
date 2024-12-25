package main

import "fmt"

func main() {
	targetWarna := []string{"merah", "kuning", "hijau", "ungu"}
	berhasil := true

	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, ": ")
		var warna1, warna2, warna3, warna4 string
		fmt.Scan(&warna1, &warna2, &warna3, &warna4)

		if warna1 != targetWarna[0] || warna2 != targetWarna[1] || warna3 != targetWarna[2] || warna4 != targetWarna[3] {
			berhasil = false
		}
	}

	fmt.Println("BERHASIL:", berhasil)
}