package main

import (
	"fmt"
	"math"
)

func main() {
	var jejari int

	fmt.Print("jejari: ")
	fmt.Scan(&jejari)

	volume := (4.0 / 3.0) * math.Pi * math.Pow(float64(jejari), 3)
	luasKulit := 4 * math.Pi * math.Pow(float64(jejari), 2)

	fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f\n", jejari, volume, luasKulit)
}
