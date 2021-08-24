package main

import (
	"fmt"
	"methods_demo/collections"
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

	fmt.Println("Products")
	fmt.Println("================================================================")

	products := collections.Products{
		{105, "Pen", 50, 5, "Stationary"},
		{107, "Pencil", 100, 2, "Stationary"},
		{103, "Marker", 20, 50, "Utencil"},
		{102, "Stove", 5, 5000, "Utencil"},
		{101, "Kettle", 10, 2500, "Utencil"},
		{104, "Scribble Pad", 20, 20, "Stationary"},
	}

	marker := models.Product{103, "Marker", 50, 20, "Utencil"}
	fmt.Println(products.Format())

	fmt.Println("IndexOf")
	fmt.Println("Index Of marker = ", products.IndexOf(marker))

	fmt.Println()
	fmt.Println("Includes")
	fmt.Println("Includes marker = ", products.Includes(marker))

	fmt.Println()
	fmt.Println("Filter")

	fmt.Println("All costly products")

	costlyProducts := products.Filter(func(p models.Product) bool {
		return p.Cost > 100
	})
	fmt.Println(costlyProducts.Format())

	fmt.Println("All stationary products")
	var stationaryProductPredicate func(models.Product) bool = func(p models.Product) bool {
		return p.Category == "Stationary"
	}
	stationaryProducts := products.Filter(stationaryProductPredicate)
	fmt.Println(stationaryProducts.Format())
}
