package main

import (
	"fmt"
)

// Struct Car
type Car struct {
	carType     string
	fuelin      float64
	consumption float64
}

// Method untuk menghitung perkiraan jarak yang bisa ditempuh
func (c Car) estimateDistance() float64 {
	estimatedDistance := c.fuelin * c.consumption
	return estimatedDistance
}

func main() {
	// Membuat objek Car
	var car Car
	fmt.Print("Enter car type: ")
	fmt.Scan(&car.carType)
	fmt.Print("Enter amount of fuel (liters): ")
	fmt.Scan(&car.fuelin)
	fmt.Print("Enter fuel consumption (L/km): ")
	fmt.Scan(&car.consumption)

	// Memanggil method estimateDistance()
	estimatedDistance := car.estimateDistance()

	// Menampilkan hasil dalam format yang diinginkan
	fmt.Printf("car type: %s , est. distance: %.2f\n", car.carType, estimatedDistance)
}
