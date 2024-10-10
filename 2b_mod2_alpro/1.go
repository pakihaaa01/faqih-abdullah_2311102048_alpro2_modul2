package main

import "fmt"

func main() {
	var eksperimen [5][4]string

	// Masukkan warna untuk setiap eksperimen
	for i := 0; i < 5; i++ {
		fmt.Printf("Masukkan warna untuk eksperimen %d:\n", i+1)
		for j := 0; j < 4; j++ {
			fmt.Printf("Masukkan warna untuk gelas %d: ", j+1)
			fmt.Scanln(&eksperimen[i][j])
		}
	}

	// Periksa apakah semua eksperimen memiliki urutan warna yang benar
	urutanBenar := true
	for _, eksperimen := range eksperimen {
		if eksperimen[0] != "merah" || eksperimen[1] != "kuning" || eksperimen[2] != "hijau" || eksperimen[3] != "ungu" {
			urutanBenar = false
			break
		}
	}

	// Cetak hasilnya
	if urutanBenar {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}
