package main

import (
	"fmt"
)

func Substring(strA, strB string) string {
	var CommonSubstring string

	for i := 0; i < len(strA); i++ {
		for j := 0; j < len(strB); j++ {
			k := 0
			for i+k < len(strA) && j+k < len(strB) && strA[i+k] == strB[j+k] {
				k++
			}
			if k > len(CommonSubstring) {
				CommonSubstring = strA[i : i+k]
			}
		}
	}
	return CommonSubstring
}

func main() {
	var stringA, stringB string

	fmt.Print("Masukkan string A: ")
	fmt.Scanln(&stringA)

	fmt.Print("Masukkan string B: ")
	fmt.Scanln(&stringB)

	commonsSubstring := Substring(stringA, stringB)

	if commonsSubstring != "" {
		fmt.Printf("Substring yang sama di antara kedua string: %s\n", commonsSubstring)
	} else {
		fmt.Println("Tidak ada substring yang sama di antara kedua string.")
	}
}
