package main

import (
	"fmt"
)

func findMaxAndMin(numbers []int) (int, int) {
	max := numbers[0]
	min := numbers[0]

	for _, num := range numbers {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return max, min
}

func main() {
	var inputNumbers [6]int

	fmt.Println("Masukkan 6 angka:")

	for i := 0; i < 6; i++ {
		fmt.Scanln(&inputNumbers[i])
	}

	maximum, minimum := findMaxAndMin(inputNumbers[:])

	maxPtr := &maximum
	minPtr := &minimum

	fmt.Printf("Nilai Maksimum Adalah: %d\n", *maxPtr)
	fmt.Printf("Nilai Minimum Adalah: %d\n", *minPtr)
}
