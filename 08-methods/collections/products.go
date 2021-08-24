package collections

import (
	"fmt"
	"methods_demo/models"
)

type Products []models.Product

func (products Products) Format() string {
	result := ""
	for _, p := range products {
		result += fmt.Sprintf("%v\n", p.Format())
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
