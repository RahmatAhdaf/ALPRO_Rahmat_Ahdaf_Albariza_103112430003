package main
import "fmt"
func main() {
	var (
		satu, dua, tiga string
 		temp string
	)
	fmt.Print("Masukan input string: ")
 	fmt.Scanln(&satu)
 	fmt.Print("Masukan input string: ")
 	fmt.Scanln(&dua)
 	fmt.Print("Masukan input string: ")
 	fmt.Scanln(&tiga)
 	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
 	temp = satu
 	satu = dua
 	dua = tiga
 	tiga = temp
 	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}

// Langkah 1: Deklarasi Variabel
// Program mendeklarasikan empat variabel string: satu, dua, tiga, dan temp. Variabel temp digunakan sebagai penyimpanan sementara untuk menggantikan nilai dari ketiga variabel lainnya. 

// Langkah 2: Input 
// Program meminta pengguna untuk memasukkan tiga nilai string menggunakan fmt.Print dan fmt.Scanln. Nilai input disimpan dalam variabel satu, dua, dan tiga, secara berurutan. 

// Langkah 3: Output Awal 
// Program mencetak nilai awal dari satu, dua, dan tiga menggunakan fmt.Println. Outputnya berformat "Output awal = <satu><dua><tiga> ". 

// Langkah 4: Penggantian Nilai 
// Program menggantikan nilai dari satu, dua, dan tiga menggunakan variabel temp. Penggantian dilakukan dalam urutan berikut: 
// •	temp mengambil nilai dari satu 
// •	satu mengambil nilai dari dua 
// •	dua mengambil nilai dari tiga 
// •	tiga mengambil nilai dari temp (yang awalnya satu) 

//Langkah 5: Output Akhir Program mencetak nilai akhir dari satu, dua, dan tiga menggunakan fmt.Println. Outputnya berformat "Output akhir = <satu><dua><tiga> ".
