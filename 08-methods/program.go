package main

import (
	"fmt"
	"methods_demo/models"
)

func main() {

	var product *models.Product = &models.Product{Id: 100, Name: "Pen", Units: 10, Cost: 5.00, Category: "Stationary"}
	fmt.Println(product)

	fmt.Printf("Cost of %s is %v\n", product.Name, product.Cost)
	fmt.Println("Before applying discount ")
	fmt.Println(product.Format())
	//applyDiscount(&product, 10)
	product.ApplyDiscount(10)
	fmt.Println("After applying discount ")
	fmt.Println(product.Format())

	banana := models.NewPerishableProduct(102, "Banana", 10, 2.00, "Fruit", "2 Days")
	banana.ApplyDiscount(10)
	fmt.Println(banana.Format())
}
