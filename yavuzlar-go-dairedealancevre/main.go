package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64

	fmt.Print("Dairenin yarıçapını girin: ")
	fmt.Scanln(&radius)

	area := math.Pi * math.Pow(radius, 2)
	circumference := 2 * math.Pi * radius

	fmt.Printf("Dairenin Alanı: %.2f\n", area)
	fmt.Printf("Dairenin Çevresi: %.2f\n", circumference)
}
