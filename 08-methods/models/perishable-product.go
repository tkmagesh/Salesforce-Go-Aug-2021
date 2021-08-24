package models

import "fmt"

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

func (p PerishableProduct) Format() string {
	//return fmt.Sprintf("Id = %d, Name = %s, Units = %d, Cost = %v, Category = %s, Expiry = %s", p.Id, p.Name, p.Units, p.Cost, p.Category, p.Expiry)
	return fmt.Sprintf("%v, Expiry = %s", p.Product, p.Expiry)
}
