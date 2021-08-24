package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

//utility functions

type ShapeWithArea interface {
	Area() float64
}

func printArea(shapeWithArea ShapeWithArea) {
	fmt.Println("Area = ", shapeWithArea.Area())
}

type ShapeWithPerimeter interface {
	Perimeter() float64
}

func printPerimeter(shapeWithPerimeter ShapeWithPerimeter) {
	fmt.Println("Perimeter = ", shapeWithPerimeter.Perimeter())
}

//interface composition
type Dimension interface {
	ShapeWithArea
	ShapeWithPerimeter
}

func printDimentions(d Dimension) {
	fmt.Println("Area = ", d.Area())
	fmt.Println("Perimeter = ", d.Perimeter())
}

func main() {
	c := Circle{10}
	/*
		printArea(c)
		printPerimeter(c)
	*/
	printDimentions(c)
	r := Rectangle{10, 20}
	/*
		printArea(r)
		printPerimeter(r)
	*/
	printDimentions(r)

}
