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

	var product *Product = &Product{Id: 100, Name: "Pen", Units: 10, Cost: 5.00, Category: "Stationary"}
	fmt.Println(product)

	fmt.Printf("Cost of %s is %v\n", product.Name, product.Cost)
	fmt.Println("Before applying discount ")
	fmt.Println(product.Format())
	//applyDiscount(&product, 10)
	product.applyDiscount(10)
	fmt.Println("After applying discount ")
	fmt.Println(product.Format())
}

func (p *Product) applyDiscount(discount int) {
	p.Cost = p.Cost * ((100 - float64(discount)) / 100)
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %s, Units = %d, Cost = %v, Category = %s", p.Id, p.Name, p.Units, p.Cost, p.Category)
}
