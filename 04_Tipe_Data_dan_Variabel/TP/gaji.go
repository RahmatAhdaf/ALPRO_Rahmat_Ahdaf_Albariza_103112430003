package main

import "fmt"

func main() {
	// mendeklarasikan variabel "jamPerMinggu" dan "upahPerJam" dengan tipe data "float64"
	var jamPerMinggu float64
	var upahPerJam float64
	// mendeklarasikan konstanta "jamNormal" adalah 40
	const jamNormal = 40
	// mendeklarasikan konstanta "lembur" adalah 1.5
	const lembur = 1.5

	// Meminta input jumlah jam kerja dalam seminggu dan upah per jam
	fmt.Print("Masukkan jumlah jam kerja per minggu: ")
	fmt.Scanln(&jamPerMinggu)
	fmt.Print("Masukkan upah per jam: ")
	fmt.Scanln(&upahPerJam)

	// mendeklarasikan variabel "bayaranNormal" dan "bayaranLembur" dengan tipe data "float"
	var bayaranNormal, bayaranLembur float64

	// Jika jam kerja lebih dari 40 jam, hitung lembur
	if jamPerMinggu > jamNormal {
		bayaranNormal = jamNormal * upahPerJam
		jamLembur := jamPerMinggu - jamNormal
		bayaranLembur = jamLembur * upahPerJam * lembur
	} else {
		bayaranNormal = jamPerMinggu * upahPerJam
		bayaranLembur = 0
	}

	// Menghitung total gaji mingguan
	totalGajiMingguan := bayaranNormal + bayaranLembur

	// Menghitung gaji bulanan (4 minggu)
	totalGajiBulanan := totalGajiMingguan * 4

	// Menampilkan total gaji bulanan
	fmt.Printf("Total gaji bulanan: %.2f\n", totalGajiBulanan)
}
