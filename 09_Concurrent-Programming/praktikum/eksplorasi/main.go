package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Product struct {
	Title        string `json:"title"`
	Price        float64 `json:"price"`
	Category     string `json:"category"`
}

func main() {
	productChannel := make(chan Product)
	
	var wg sync.WaitGroup

	resp, err := http.Get("https://fakestoreapi.com/products")
	if err != nil {
		fmt.Println("Gagal mengambil data:", err)
		return
	}
	defer resp.Body.Close()

	var products []Product
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		fmt.Println("Gagal menguraikan respons JSON:", err)
		return
	}

	for _, p := range products {
		wg.Add(1)
		go func(product Product) {
			defer wg.Done()
			productChannel <- product
		}(p)
	}

	go func() {
		wg.Wait()
		close(productChannel)
	}()

	for product := range productChannel {
		fmt.Printf("Title: %s\n", product.Title)
		fmt.Printf("Price: $%.2f\n", product.Price)
		fmt.Printf("Category: %s\n", product.Category)
		fmt.Println("==============================")
	}
}
