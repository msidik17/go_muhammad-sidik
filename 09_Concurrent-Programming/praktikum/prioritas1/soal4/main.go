package main

import (
	"fmt"
	"sync"
)

func fibonacci(n int) int {
	var mutex sync.Mutex
	if n <= 0 {
		return 0
	}

	fib := make([]int, n+1)
	fib[0] = 0
	fib[1] = 1
	
	mutex.Lock()
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	defer mutex.Unlock()
	return fib[n]
}

func main() {
	var n int
	fmt.Print("Masukkan indeks F(n): ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Input tidak valid")
		return
	}
	
	result := fibonacci(n)
	fmt.Printf("F(%d) = %d\n", n, result)
}
