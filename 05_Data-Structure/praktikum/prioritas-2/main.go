package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cari (angka []int, target int)[]int {
	left, right := 0, len(angka)-1

	for left < right {
		jumlah := angka[left] + angka[right]

		if jumlah == target {
			return []int{left, right}
		}else if jumlah < target {
			left++
		}else{
			right--
		}
	}
	return nil
}

func main(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Data Array: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var target int
	fmt.Print("Masukkan Target: ")
	fmt.Scanln(&target)

	inputAngka := strings.Split(input,",")
	var angka []int
	for _, numStr := range inputAngka{
		var num int
		fmt.Sscanf(numStr, "%d", &num)
		angka = append(angka, num)
	}

	hasil := cari(angka, target)
	if hasil != nil{
		fmt.Printf("Angka Berpasangan Adalah : [%d,  %d]\n",hasil[0], hasil[1])
	}else {
		fmt.Printf("Tidak Menemukan Pasangan Yang Sesuai")
	}
}

