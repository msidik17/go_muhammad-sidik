package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	var result []pair // Mengganti "[]pairs" dengan "[]pair"
	itemCount := make(map[string]int)
	for _, item := range items {
		itemCount[item]++
	}

	for key, count := range itemCount {
		result = append(result, pair{name: key, count: count})
	}

	// Mengurutkan hasil berdasarkan count
	sort.Slice(result, func(i, j int) bool {
		return result[i].count > result[j].count
	})

	return result
}

func main() {
	pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}) // golang->1 ruby->2 js->4
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}) // C->1 D->2 B->3 A->4
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"football", "basketball", "tenis"}) // football->1 basketball->1 tenis->1
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()
}
