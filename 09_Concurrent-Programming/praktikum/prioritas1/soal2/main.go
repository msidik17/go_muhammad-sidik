package main

import (
	"fmt"
	"time"
)

func main(){

x := 3
multipleTime := 5
resultTemp := 0

tempMultiple := make(chan int, multipleTime)
func() {
	
	for i := 0; i < multipleTime; i++ {
		resultTemp += x
		tempMultiple <-resultTemp
		time.Sleep(1 * time.Second)
	}
	
}()
go func() {
	
	for i := 0; i < multipleTime; i++ {
		result := <-tempMultiple
		fmt.Println(result)
		time.Sleep(1 * time.Second)
	}
	
}()

	time.Sleep(time.Duration(multipleTime) * 1 * time.Second)
}