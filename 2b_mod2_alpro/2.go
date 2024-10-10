package main

import "fmt"

func main() {
	var n int
	fmt.Print("N: ")
	fmt.Scanln(&n)

	var pita string
	for i := 0; i < n; i++ {
		var namaBunga string
		fmt.Printf("Bunga %d: ", i+1)
		fmt.Scanln(&namaBunga)

		if i > 0 {
			pita += " - "
		}
		pita += namaBunga
	}

	fmt.Println("Pita:", pita)
}
