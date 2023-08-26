package main

import (
	"fmt"
	"math"
)

func main(){
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	n := len(matrix)
	
	fmt.Println("Matriks yang dimasukkan:")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

	sumLeft := 0
	for i := 0; i < n; i++ {
		sumLeft += matrix[i][i]
	}

	sumRight := 0
	for i := 0; i < n; i++ {
		sumRight += matrix[i][n-i-1]
	}

	diff := int(math.Abs(float64(sumLeft - sumRight)))

	fmt.Printf("Perbedaan Mutlak adalah : %d\n", diff)
}