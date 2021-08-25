package collections

import (
	"fmt"
	"methods_demo/models"
	"sort"
)

type Products []models.Product

//implementing the fmt.Stringer interface
func (products Products) String() string {
	result := ""
	for _, p := range products {
		result += fmt.Sprintf("%v\n", p)
	}
	return result
}

func (products Products) IndexOf(product models.Product) int {
	for i, p := range products {
		if p == product {
			return i
		}
	}
	return -1
}

func (products Products) Includes(product models.Product) bool {
	for _, p := range products {
		if p == product {
			return true
		}
	}
	return false
}

func (products Products) Filter(predicate func(models.Product) bool) Products {
	var result Products
	for _, p := range products {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

func (products Products) Any(predicate func(models.Product) bool) bool {
	for _, p := range products {
		if predicate(p) {
			return true
		}
	}
	return false
}

func (products Products) All(predicate func(models.Product) bool) bool {
	for _, p := range products {
		if !predicate(p) {
			return false
		}
	}
	return true
}

//Sorting

//"sort.Interface" implementation

func (products Products) Len() int {
	return len(products)
}

//compare the products by id
func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

func (products Products) Sort() {
	sort.Sort(products)
}

//to sort by name
type byName struct {
	Products
}

func (products byName) Less(i, j int) bool {
	return products.Products[i].Name < products.Products[j].Name
}

func (products Products) SortByName() {
	sort.Sort(byName{products})
}

//to sort by cost
func (products Products) SortByCost() {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Cost < products[j].Cost
	})
}
