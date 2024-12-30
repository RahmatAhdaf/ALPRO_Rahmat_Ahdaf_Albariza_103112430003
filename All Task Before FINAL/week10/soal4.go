// PSEUDOCODE
// 
// program DurasiParkir
// kamus
//     h1, m1, h2, m2 : integer
//     durasiJam, durasiMenit : integer
// 
// algoritma
//     input(h1, m1, h2, m2)
//     if (h2 < h1) or (h2 = h1 and m2 < m1) then
//         h2 ← h2 + 12
//     endif
//     durasiJam ← h2 - h1
//     durasiMenit ← m2 - m1
//     if durasiMenit < 0 then
//         durasiMenit ← durasiMenit + 60
//         durasiJam ← durasiJam - 1
//     endif
//     output(durasiJam, durasiMenit)
// endprogram


// CODING
package main

import "fmt"

func main() {
	// Input jam dan menit mulai serta selesai
	var h1, m1, h2, m2 int
	fmt.Print("Masukkan jam dan menit mulai parkir: ")
	fmt.Scan(&h1, &m1)
	fmt.Print("Masukkan jam dan menit selesai parkir: ")
	fmt.Scan(&h2, &m2)

	// Periksa jika waktu selesai lebih kecil dari waktu mulai (format 12 jam)
	if h2 < h1 || (h2 == h1 && m2 < m1) {
		h2 += 12
	}

	// Hitung durasi parkir
	durasiJam := h2 - h1
	durasiMenit := m2 - m1
	if durasiMenit < 0 {
		durasiMenit += 60
		durasiJam--
	}

	// Output hasil
	fmt.Printf("Durasi parkir: %d jam %d menit\n", durasiJam, durasiMenit)
}
