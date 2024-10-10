package main

import "fmt"

func isKabisat(tahun int) bool {
	if tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0) {
		return true
	}
	return false
}

func main() {
	var tahun int
	fmt.Print("Masukkan tahun: ")
	fmt.Scanln(&tahun)

	kabisat := isKabisat(tahun)
	fmt.Printf("Tahun: %d\nKabisat: %t\n", tahun, kabisat)
}
