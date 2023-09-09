package main

import (
	"fmt"
)

func main(){

bufferChan := make(chan int, 5)

	go multiple(bufferChan)

	for i := 0; i < 5; i++ {
		resultTemp := <-bufferChan
		fmt.Println("Bilangan Kelipatan 3 :", resultTemp)		
	}
}

func multiple ( bufferChan chan int) {
	
	for i := 1; i <= 5; i++ {
		bufferChan <-i * 3
	}
	close(bufferChan)
}