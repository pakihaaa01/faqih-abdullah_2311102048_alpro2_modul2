package main

import "fmt"

func main() {
	var bil int

	// Meminta input dari pengguna
	fmt.Print("Bilangan: ")
	fmt.Scan(&bil)

	// Menampilkan faktor-faktor bilangan
	fmt.Print("Faktor: ")
	for i := 1; i <= bil; i++ {
		if bil%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
