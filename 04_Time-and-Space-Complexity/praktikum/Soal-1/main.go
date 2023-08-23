package main

import (
	"fmt"
	"time"
)

func linier(n int)bool{
	if n <= 1{
		return false
	} 
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main(){
	var n int 
	fmt.Printf("Masukkan Bilangan :")
	fmt.Scanln(&n)

	if linier(n){
		fmt.Printf("%d Adalah Bilangan Prima.  \n", n)
	}else{
		fmt.Printf("%d Bukan Bilangan Prima.  \n", n)
	}

	endtime := time.Now()
	fmt.Println(endtime)
}