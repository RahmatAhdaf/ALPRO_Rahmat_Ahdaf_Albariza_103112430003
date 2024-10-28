package main

import "fmt"

// Bentukan untuk menyimpan data karyawan
type Karyawan struct {
	Nama      string
	GajiPokok float64
	Tunjangan float64
	Potongan  float64
}

// Fungsi untuk menghitung total gaji karyawan
func (k Karyawan) TotalGaji() float64 {
	return k.GajiPokok + k.Tunjangan - k.Potongan
}

func main() {
	var nama string
	var gajiPokok, tunjangan, potongan float64

	// Masukan data karyawan
	fmt.Println("Masukkan nama karyawan:")
	fmt.Scanln(&nama)
	fmt.Println("Masukkan gaji pokok:")
	fmt.Scanln(&gajiPokok)
	fmt.Println("Masukkan tunjangan:")
	fmt.Scanln(&tunjangan)
	fmt.Println("Masukkan potongan:")
	fmt.Scanln(&potongan)

	// Simpan data ke struct Karyawan
	karyawan := Karyawan{
		Nama:      nama,
		GajiPokok: gajiPokok,
		Tunjangan: tunjangan,
		Potongan:  potongan,
	}

	// Hitung dan cetak total gaji
	totalGaji := karyawan.TotalGaji()
	fmt.Printf("\nInformasi Karyawan:\n")
	fmt.Printf("Nama: %s\n", karyawan.Nama)
	fmt.Printf("Gaji Pokok: Rp %.2f\n", karyawan.GajiPokok)
	fmt.Printf("Tunjangan: Rp %.2f\n", karyawan.Tunjangan)
	fmt.Printf("Potongan: Rp %.2f\n", karyawan.Potongan)
	fmt.Printf("Total Gaji: Rp %.2f\n", totalGaji)
}
