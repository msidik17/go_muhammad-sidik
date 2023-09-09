package main

import (
	"fmt"
	"time"
)

func main(){

var x int
multipleTime := 5
result := 0

fmt.Print("Masukkan Angka :")
fmt.Scan(&x)

go func() {
	
	for i := 0; i < multipleTime; i++ {
		result += x
		fmt.Println(result)
		time.Sleep(3 * time.Second)
	}
	
}()
	time.Sleep(time.Duration(multipleTime) * 3 * time.Second)
}