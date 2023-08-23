package main

import (
	"fmt"
	"time"
)

func angka (x, n int)int{
	if n == 1 {
		return x
	}else if n%2 == 0 {
		kuadrat := angka(x, n/2)
		return kuadrat * kuadrat
	}else {
		kuadrat := angka (x, (n-1)/2)
		return x * kuadrat * kuadrat 
	}
}

func main() {
	var (x, n int)

	fmt.Printf("Masukkan Nilai X : ")
	fmt.Scanln(&x)

	fmt.Printf("Masukkan Nilai N : ")
	fmt.Scanln(&n)

	result := angka(x, n)
	fmt.Println("Hasil : ", result)

	endtime := time.Now()
	fmt.Println(endtime)


}