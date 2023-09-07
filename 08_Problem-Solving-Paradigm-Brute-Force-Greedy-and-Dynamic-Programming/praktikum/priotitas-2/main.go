package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func frog(jumps []int) int {
	n := len(jumps)
	x := make([]int, n)

	x[0] = 0 
	x[1] = abs(jumps[1] - jumps[0])

	for i := 2; i < n; i++ {
		jump1 := x[i-1] + abs(jumps[i] - jumps[i-1])
		jump2 := x[i-2] + abs(jumps[i] - jumps[i-2])
		x[i] = min(jump1, jump2)
	}
	return x[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Print("Input  : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Membagi input menjadi slice berdasarkan spasi
	inputArr := strings.Fields(input)

	var jumps []int

	for _, str := range inputArr {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Input tidak valid.")
			return
		}
		jumps = append(jumps, num)
	}

	if len(jumps) < 2 {
		fmt.Println("Tidak cukup batu untuk melompat.")
		return
	}

	minCost := frog(jumps)
	fmt.Printf("Output : %d\n", minCost)
}