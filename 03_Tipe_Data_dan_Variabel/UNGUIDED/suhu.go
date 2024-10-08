package main

import "fmt"

func main() {
	// mendeklarasikan variabel "celsius" dengan tipe data "float64"
    var celsius float64

	// manampilkan pesan yang menyuruh pengguna menginputkan temperatur celsius
    fmt.Print("Suhu Celsius: ")
    fmt.Scan(&celsius)

    // rumus konversi dari Celsius ke Reamur, Fahrenheit, dan Kelvin
    reamur := celsius * 4 / 5
    fahrenheit := (celsius * 9 / 5) + 32
    kelvin := celsius + 273.15

    // menampilkan hasil konversi
    fmt.Printf("Suhu Reamur: %.2f\n", reamur)
    fmt.Printf("Suhu Fahrenheit: %.2f\n", fahrenheit)
    fmt.Printf("Suhu Kelvin: %.2f\n", kelvin)
}