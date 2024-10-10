package main

import "fmt"

func hitungBiaya(beratGram int) (int, int, int) {
	// Menghitung berat dalam kg dan sisa gram
	beratKg := beratGram / 1000
	sisaGram := beratGram % 1000

	// Biaya dasar per kg
	biayaDasar := beratKg * 10000

	// Perhitungan biaya tambahan berdasarkan sisa gram
	biayaTambahan := 0
	if beratKg > 10 {
		// Jika total berat lebih dari 10kg, sisa gram digratiskan
		biayaTambahan = 0
	} else if sisaGram >= 500 {
		// Jika sisa gram >= 500, tambahan biaya Rp 5 per gram
		biayaTambahan = sisaGram * 5
	} else {
		// Jika sisa gram < 500, tambahan biaya Rp 15 per gram
		biayaTambahan = sisaGram * 15
	}

	// Total biaya
	totalBiaya := biayaDasar + biayaTambahan

	return beratKg, sisaGram, totalBiaya
}

func main() {
	var beratParsel int
	var beratKg, sisaGram, totalBiaya, biayaTambahan int

	// Meminta input dari pengguna
	fmt.Print("Masukkan berat parsel (dalam gram): ")
	fmt.Scan(&beratParsel)

	// Menghitung biaya pengiriman
	beratKg, sisaGram, totalBiaya = hitungBiaya(beratParsel)

	// Menampilkan hasil
	fmt.Println("Berat parsel (gram):", beratParsel)
	fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", beratKg*10000, biayaTambahan)
	fmt.Println("Total biaya: Rp.", totalBiaya)
}
