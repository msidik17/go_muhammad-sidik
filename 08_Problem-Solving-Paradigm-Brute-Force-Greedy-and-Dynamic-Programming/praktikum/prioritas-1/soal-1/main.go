package main

import (
	"fmt"
	"strconv"
)

func binary(n int) []string {
	ans := make([]string, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = strconv.FormatInt(int64(i), 2)
	}
	return ans
}

func main() {
	var n int

	fmt.Print("Masukkan Bilangan          :" )
	_, err := fmt.Scan(&n)
	if err != nil{
		fmt.Println("Input Bilangan Tidak Valid", err)
		return
	}
	ans := binary(n)
	fmt.Println("Representasi Binary Adalah :",ans)
}
