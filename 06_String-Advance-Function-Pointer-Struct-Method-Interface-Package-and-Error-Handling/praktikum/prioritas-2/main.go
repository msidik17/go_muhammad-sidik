package main

import (
	"fmt"
)

func caesarCipher(offset int, input string) string {
	result := ""
	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			shifted := 'a' + (char-'a'+rune(offset))%26
			result += string(shifted)
		} else if char == ' ' {
			result += " "
		}
	}
	return result
}

func main() {
	var offset int
	var input string

	fmt.Print("Masukkan offset: ")
	fmt.Scan(&offset)

	fmt.Print("Masukkan input: ")
	fmt.Scan(&input)

	output := caesarCipher(offset, input)
	fmt.Printf("Hasil enkripsi: %s\n", output)
}
