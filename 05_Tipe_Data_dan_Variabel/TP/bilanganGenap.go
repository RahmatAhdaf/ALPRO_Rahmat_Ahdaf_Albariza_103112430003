package main

import "fmt"

func main() {
    // Menampilkan bilangan genap dari 1 hingga 50
    fmt.Println("Bilangan genap dari 1 hingga 50:")
    for i := 1; i <= 50; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }
}
