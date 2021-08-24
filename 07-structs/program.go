package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Units    int
	Cost     float64
	Category string
}

func main() {
	/*
		var product Product
		product.Id = 100
		product.Name = "Pen"
		product.Units = 10
		product.Cost = 5.00
		product.Category = "Stationary"
	*/
	//var product Product = Product{100, "Pen", 10, 5.00, "Stationary"}
	var product Product = Product{Id: 100, Name: "Pen", Units: 10, Cost: 5.00, Category: "Stationary"}
	fmt.Println(product)

	fmt.Printf("Cost of %s is %v\n", product.Name, product.Cost)
	fmt.Println("Before applying discount ", product)
	applyDiscount(&product, 10)
	fmt.Println("After applying discount ", product)

	/*
		var p2 = new(Product)
		fmt.Println(p2)
	*/
}

func applyDiscount(p *Product, discount int) {
	//(*p).Cost = (*p).Cost * ((100 - float64(discount)) / 100)
	p.Cost = p.Cost * ((100 - float64(discount)) / 100)
}
