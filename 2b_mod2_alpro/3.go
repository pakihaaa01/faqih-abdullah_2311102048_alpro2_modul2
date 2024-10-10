package main

import "fmt"

func main() {
	var beratKantong1, beratKantong2 float64
	var totalBerat float64

	for {
		fmt.Printf("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&beratKantong1, &beratKantong2)

		// Cek kondisi berhenti
		totalBerat = beratKantong1 + beratKantong2
		if totalBerat >= 150 || beratKantong1 < 0 || beratKantong2 < 0 {
			break
		}

		// Cek kondisi oleng
		selisihBerat := abs(beratKantong1 - beratKantong2)
		if selisihBerat >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}

	fmt.Println("Proses selesai.")
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
