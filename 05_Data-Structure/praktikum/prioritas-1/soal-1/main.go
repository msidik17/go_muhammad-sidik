package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func text (array []string)[]string {
	cari := map[string]bool{}
	hasil := []string {}

	for _, n := range array{
		if !cari[n] {
			cari[n] = true
			hasil = append(hasil, n)
		}
	}
	return hasil
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Masukkan Data Array 1 :")
	array1, _ := reader.ReadString('\n')

	fmt.Print("Masukkan Data Array 2 :")
	array2, _ := reader.ReadString('\n')

	data1 := strings.Split(strings.TrimSpace(array1), ",")
	data2 := strings.Split(strings.TrimSpace(array2), ",")

	combine := append(data1, data2...)
	hasilcombine := text(combine)

	fmt.Println(hasilcombine)

}