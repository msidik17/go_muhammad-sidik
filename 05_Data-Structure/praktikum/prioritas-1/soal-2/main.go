package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Data Array  :")
	array, _ := reader.ReadString('\n')

	data := strings.Split(strings.TrimSpace(array), ",")

	hitung := make(map[string]int)

	for _, nilai := range data {
		hitung[nilai]++
	}

	var output []string
	for nilai, hitung := range hitung{
		output = append(output, fmt.Sprintf("%s : %d", nilai, hitung))
	}

	sort.Strings(output)
	
	fmt.Printf("[%s]\n", strings.Join(output, "  "))
}