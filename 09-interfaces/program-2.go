package main

import "fmt"

type Product struct {
	Name string
}

func (p Product) Format() {
}

type Formatter interface {
	Format()
}

func main() {
	var x interface{}
	x = 100
	/*
		x = "This is a string"
		x = true
		x = struct{}{}
		x = []int{10, 20, 30}
		fmt.Println(x)
	*/

	if no, ok := x.(int); ok {
		y := no + 200
		fmt.Println(y)
	} else {
		fmt.Println("Not an int")
	}

	//x = "This is a string"
	//x = struct{}{}
	//x = Product{Name: "Pen"}
	switch val := x.(type) {
	case int, float64:
		fmt.Println("x is either an int or float64")
	case string:
		fmt.Println("Length of the x is ", len(val))
	case bool:
		fmt.Println("x is ", val)
	case struct{}:
		fmt.Println("x is a struct")
	case []int:
		fmt.Println("x is a slice of int")
	case Product:
		fmt.Println("x is a Product")
	case Formatter:
		fmt.Println("x implements Formatter interface ")
	default:
		fmt.Println("Unknown type")
	}

	if _, ok := x.(Formatter); ok {
		fmt.Println(" x is a Formatter")
	} else {
		fmt.Println("Not an int")
	}

}
