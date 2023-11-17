package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Birinci sayıyı girin: ")
	fmt.Scanln(&num1)

	fmt.Print("İkinci sayıyı girin: ")
	fmt.Scanln(&num2)

	fmt.Print("İşlemi seçin (+, -, *, /): ")
	fmt.Scanln(&operator)

	result := 0.0

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Hata: Sıfıra bölme hatası")
			return
		}
	default:
		fmt.Println("Hata: Geçersiz işlem")
		return
	}

	fmt.Printf("Sonuç: %s %s %s = %f\n", strconv.FormatFloat(num1, 'f', -1, 64), operator, strconv.FormatFloat(num2, 'f', -1, 64), result)
}
