package main

import "fmt"

func SimpleEquations(a, b, c int) (int, int, int) {
	for x := 1; x <= a; x++ {
		for y := 1; y <= a; y++ {
			z := a - x - y
			if x*y*z == b && x*x+y*y+z*z == c {
				return x, y, z
			}
		}
	}
	return 0, 0, 0 
}


func main() {
	var A, B, C int
	fmt.Print("Input : ")
	fmt.Scan(&A, &B, &C)

	x, y, z := SimpleEquations(A, B, C)
	if x == 0 && y == 0 && z == 0 {
		fmt.Println("No Solution")
	} else {
		fmt.Printf("Combination : %d %d %d\n", x, y, z)
	}
}