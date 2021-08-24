package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func main() {
	products := []Product{
		{105, "Pen", 5, 50, "Stationary"},
		{107, "Pencil", 2, 100, "Stationary"},
		{103, "Marker", 50, 20, "Utencil"},
		{102, "Stove", 5000, 5, "Utencil"},
		{101, "Kettle", 2500, 10, "Utencil"},
		{104, "Scribble Pad", 20, 20, "Stationary"},
	}

	marker := Product{103, "Marker", 50, 20, "Utencil"}
	fmt.Println(FormatProducts(products))

	fmt.Println("IndexOf")
	fmt.Println("Index Of marker = ", IndexOf(marker, products))

	fmt.Println()
	fmt.Println("Includes")
	fmt.Println("Includes marker = ", Includes(marker, products))

	fmt.Println()
	fmt.Println("Filter")
	/*
		fmt.Println("All costly products")
		costlyProducts := FilterCostlyProducts(products)
		fmt.Println(FormatProducts(costlyProducts))

		fmt.Println("All stationary products")
		stationaryProducts := FilterStationaryProducts(products)
		fmt.Println(FormatProducts(stationaryProducts))
	*/

	fmt.Println("All costly products")
	/*
		var costlyProductPredicate func(Product) bool = func(p Product) bool {
			return p.Cost > 100
		}
		costlyProducts := Filter(products, costlyProductPredicate)
	*/
	costlyProducts := Filter(products, func(p Product) bool {
		return p.Cost > 100
	})
	fmt.Println(FormatProducts(costlyProducts))

	fmt.Println("All stationary products")
	var stationaryProductPredicate func(Product) bool = func(p Product) bool {
		return p.Category == "Stationary"
	}
	stationaryProducts := Filter(products, stationaryProductPredicate)
	fmt.Println(FormatProducts(stationaryProducts))

}

func Format(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %s, Units = %d, Cost = %v, Category = %s", p.Id, p.Name, p.Units, p.Cost, p.Category)
}

func FormatProducts(products []Product) string {
	result := ""
	for _, p := range products {
		result += fmt.Sprintf("%v\n", Format(p))
	}
	return result
}

//IndexOf => returns the index of the given product in the products collections (-1 if it doesnt exist)

func IndexOf(product Product, products []Product) int {
	for i, p := range products {
		if p == product {
			return i
		}
	}
	return -1
}

//Includes => return true/false based on the existence of the given product in the products collection

func Includes(product Product, products []Product) bool {
	for _, p := range products {
		if p == product {
			return true
		}
	}
	return false
}

/*
Filter => returns a new collection of products that match the given criteria
use cases
filter costly products (cost > 100)
filter products by category (category == "Stationary")
filter products by category and cost (category == "Stationary" && cost > 100)
*/

/*
func FilterCostlyProducts(products []Product) []Product {
	var costlyProducts []Product
	for _, p := range products {
		if p.Cost > 100 {
			costlyProducts = append(costlyProducts, p)
		}
	}
	return costlyProducts
}

func FilterStationaryProducts(products []Product) []Product {
	var stationaryProducts []Product
	for _, p := range products {
		if p.Category == "Stationary" {
			stationaryProducts = append(stationaryProducts, p)
		}
	}
	return stationaryProducts
}
*/

func Filter(products []Product, predicate func(Product) bool) []Product {
	var result []Product
	for _, p := range products {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

/*
Any => return true/false based on the existence of ANY product in the collection that satisfies the given criteria
use cases
Are there any costly products (cost > 100)
Are there any "Stationary" products
*/

func Any(products []Product, predicate func(Product) bool) bool {
	for _, p := range products {
		if predicate(p) {
			return true
		}
	}
	return false
}

/*
All => return true/false if ALL the products in the collection satisfies the given criteria
use cases
Are all the products costly products (cost > 100)
Are all the products "Stationary" products
*/

func All(products []Product, predicate func(Product) bool) bool {
	for _, p := range products {
		if !predicate(p) {
			return false
		}
	}
	return true
}
