package main

import (
	"fmt"
)

// Fungsi untuk memeriksa apakah sebuah bilangan adalah bilangan prima atau tidak
func isPrime(num int) bool {
	if num <= 1 {
		return false
	} else if num <= 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	i := 5
	for i*i <= num {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func primeX(n int) int {
	count := 0
	num := 2 

	for {
		if isPrime(num) {
			count++
			if count == n {
				return num
			}
		}
		num++
	}
}

func main() {
	var n int
	fmt.Print("Masukkan urutan bilangan prima yang ingin ditampilkan: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Urutan bilangan prima harus lebih dari 0.")
	} else {
		result := primeX(n)
		fmt.Printf("Bilangan prima urutan ke-%d adalah: %d\n", n, result)
	}
}
