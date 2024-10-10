package main

import "fmt"

func main() {
	var celcius int

	fmt.Print("Maukkan Temperatur Celcius: ")
	fmt.Scan(&celcius)

	// Konversi ke Fahrenheit
	fahrenheit := int(float64(celcius)*9.0/5.0 + 32)

	// Konversi ke Reamur
	reamur := int(float64(celcius) * 4.0 / 5.0)

	// Konversi ke Kelvin
	kelvin := int(float64(celcius) + 273.15)

	fmt.Printf("Temperatur celsius: %d\n", celcius)
	fmt.Printf("Derajat Reamur: %d\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %d\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %d\n", kelvin)
}
