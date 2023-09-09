package main

import (
	"fmt"
)

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna"
	frequency := make(map[rune]int)

	for _, char := range text {
		frequency[char]++
	}

	for char, count := range frequency {
		fmt.Printf("%c: %d\n", char, count)
	}
}
