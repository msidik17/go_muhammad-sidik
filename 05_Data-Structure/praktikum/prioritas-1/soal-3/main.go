package main

import (
	"fmt"
	"strings"
)

func cari(array string) []int{
	angka := strings.Split(array, "")
	hitung := make(map[int]int)

	for _, n := range angka {
		var num int
		fmt.Sscanf(n, "%d", &num)
		hitung[num]++
	}


	var angkaunik []int
	for num, nilai := range hitung {
		if nilai == 1 {
			angkaunik = append(angkaunik, num)
		}
	}
	return angkaunik
}

func main(){
	var array string
	fmt.Print("Masukkan Angka: ")
	fmt.Scanln(&array)
	
	angkaunik := cari(array)
	
	fmt.Println("Angka Yang Muncul Sekali : ",angkaunik)
}

