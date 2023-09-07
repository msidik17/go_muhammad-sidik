package main

import "fmt"

func segitigaPascal(rows int) [][]int {
	generate := make([][]int, rows)
	for i := 0; i < rows; i++ {
		row := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row[j] = i
			} else {
				row[j] = generate[i-1][j-1] + generate[i-1][j]
			}
		}
		generate[i] = row
	}
	return generate
}

func main() {
	rows := 5
	result := segitigaPascal(rows)
	fmt.Println(result)
}