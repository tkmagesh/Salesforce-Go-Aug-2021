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
	Product
	Expiry string
}

func NewPerishableProduct(id int, name string, units int, cost float64, category string, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{Id: id, Name: name, Units: units, Cost: cost, Category: category},
		Expiry:  expiry,
	}
}

func main() {

	/* banana := PerishableProduct{
		Product{Id: 100, Name: "banana", Units: 5, Cost: 1.99, Category: "fruit"},
		"2 Days",
	} */

	/*
		banana := PerishableProduct{
			Product: Product{Id: 100, Name: "banana", Units: 5, Cost: 1.99, Category: "fruit"},
			Expiry:  "2 Days",
		}
	*/
	banana := NewPerishableProduct(100, "banana", 5, 1.99, "fruit", "2 Days")
	fmt.Println(banana)

	/* fmt.Printf("Cost of %s is %v\n", banana.Product.Name, banana.Product.Cost) */
	fmt.Printf("Cost of %s is %v\n", banana.Name, banana.Cost)

	fmt.Println("Before applying discount, banana => ", banana)
	applyDiscount(&banana.Product, 10)
	fmt.Println("After applying discount, banana => ", banana)

}

func applyDiscount(p *Product, discount int) {
	//(*p).Cost = (*p).Cost * ((100 - float64(discount)) / 100)
	p.Cost = p.Cost * ((100 - float64(discount)) / 100)
}
