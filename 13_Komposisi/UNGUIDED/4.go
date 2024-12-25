package main

import (
	"fmt"
)

func main() {
	var pita string
	var bungaCount int

	for {
		var bunga string
		fmt.Printf("Bunga %d: ", bungaCount+1)
		fmt.Scanln(&bunga)

		if bunga == "SELESAI" || bunga == "selesai" {
			break
		}

		if bungaCount > 0 {
			pita += " - "
		}
		pita += bunga
		bungaCount++
	}

	fmt.Printf("Pita: %s\n", pita)
	fmt.Printf("Bunga: %d\n", bungaCount)
}