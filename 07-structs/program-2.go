package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Units    int
	Cost     float64
	Category string
}

type PerishableProduct struct {
	Prod   Product
	Expiry string
}

func main() {
	/*
		banana := PerishableProduct{
			Product{Id: 100, Name: "banana", Units: 5, Cost: 1.99, Category: "fruit"},
			"2 Days",
		}
	*/
	banana := PerishableProduct{
		Prod:   Product{Id: 100, Name: "banana", Units: 5, Cost: 1.99, Category: "fruit"},
		Expiry: "2 Days",
	}
	fmt.Println(banana)

	fmt.Printf("Cost of %s is %v\n", banana.Prod.Name, banana.Prod.Cost)

}
