package main

import (
	"fmt"
)

func main() {
	var k int
	var akarDua float64

	fmt.Print("Masukkan nilai K (jumlah iterasi): ")
	fmt.Scanln(&k)

	// Inisialisasi akarDua dengan 1
	akarDua = 1.0

	// Hitung aproksimasi akarDua menggunakan rumus yang diberikan
	for i := 0; i < k; i++ {
		akarDua *= (4*float64(i) + 2) * (4*float64(i) + 2) / ((4*float64(i) + 1) * (4*float64(i) + 3))
	}

	// Format output hingga 10 angka di belakang koma
	fmt.Printf("Aproksimasi akar 2 dengan %d iterasi adalah: %.10f\n", k, akarDua)
}
