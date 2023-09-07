package main

import (
	"fmt"
	"strings"
)

var romanNumer = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func intToRoman(num int) string {
	var result strings.Builder

	keys := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for _, value := range keys {
		for num >= value {
			result.WriteString(romanNumer[value])
			num -= value
		}
	}

	return result.String()
}

func main() {
	var num int
	fmt.Print("Masukkan Angka: ")
	_, err := fmt.Scanf("%d", &num)

	if err != nil {
		fmt.Println("Input tidak valid.")
		return
	}

	roman := intToRoman(num)
	fmt.Printf("%d dalam angka Romawi adalah %s\n", num, roman)
}
