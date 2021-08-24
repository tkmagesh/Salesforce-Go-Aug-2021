package models

import "fmt"

//state
type Product struct {
	Id       int
	Name     string
	Units    int
	Cost     float64
	Category string
}

//behavior
func (p *Product) ApplyDiscount(discount int) {
	p.Cost = p.Cost * ((100 - float64(discount)) / 100)
}

//implementing the Stringer interface (which is used by the Print functions in the fmt package)
func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %s, Units = %d, Cost = %v, Category = %s", p.Id, p.Name, p.Units, p.Cost, p.Category)
}
