package main

import "fmt"


func fibonacci(n int, temp map[int]int) int {
	if n <= 0 {
		return 0
	}
	
	if n == 1 {
		return 1
	}
	
	if val, ok := temp[n]; ok {
		return val
	}
	
	temp[n] = fibonacci(n-1, temp) + fibonacci(n-2, temp)
	return temp[n]
}


func main() {
	var n int
	fmt.Print("Masukkan indeks F(n): ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Input tidak valid")
		return
	}
	
	temp := make(map[int]int)
	result := fibonacci(n, temp)
	fmt.Printf("F(%d) = %d\n", n, result)
}
