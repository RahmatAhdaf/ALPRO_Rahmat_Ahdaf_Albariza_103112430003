package main

import "fmt"

func main() {
    // mendeklarasikan variabel "d1" "d2" dan "d3" dengan tipe data "int"
    var bilangan, d1, d2, d3 int
    // mengambil input bilangan dari pengguna dan menyimpannya ke variabel bilangan
    fmt.Scan(&bilangan)
    // memisahkan digit bilangan
    d1 = bilangan / 100
    d2 = bilangan % 100 / 10
    d3 = bilangan % 100 % 10
    // menampilkan output
    fmt.Println(d1 <= d2 && d2 <= d3)
}
