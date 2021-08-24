package main

import "fmt"

func main() {
	//Array
	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	//var nos [5]int = [...]int{3, 1, 4, 2, 5}
	nos := [...]int{3, 1, 4, 2, 5, 8, 9, 10}
	fmt.Println(nos)

	//iterating a array
	for i := 0; i < len(nos); i++ {
		fmt.Println(nos[i])
	}

	fmt.Println("Iterating using range construct")
	/*
		for idx, no := range nos {
			fmt.Println(idx, no)
		}
	*/
	for _, no := range nos {
		fmt.Println(no)
	}

	//Slices
	//var products []string = []string{"Pen", "Pencil", "Marker"}
	fmt.Println("Slices")
	var products []string = make([]string, 0, 10)
	fmt.Println(products)
	fmt.Println(len(products), cap(products))

	//adding elements to slice
	products = append(products, "Pen", "Pencil", "Marker", "Eraser")
	fmt.Println(len(products), cap(products))

	products = append(products, "ScribblePad")
	fmt.Println(len(products), cap(products))

	/*
		products = append(products, "Mouse")
		fmt.Println(len(products), cap(products))

		products = append(products, "Stylus")
		fmt.Println(len(products), cap(products))
	*/

	/*
		//appending more than one item
		products = append(products, "Mouse", "Stylus")
	*/

	newProducts := []string{"Mouse", "Stylus"}
	products = append(products, newProducts...)

	fmt.Println(products)
	/* for idx := 0; idx < 12; idx++ {
		fmt.Println(products[idx])
	} */

	//slicing elements
	/*
		[lo : hi] => from lo to hi-1
		[lo : ] => from lo to end of the slice
		[: hi] => from 0 to hi-1
		[lo : lo] => []
		[lo : lo+1] => [lo]
		[:] => copy of the slice
	*/
	fmt.Println("Products => ", products)
	fmt.Println("products[1:3] => ", products[1:3])
	fmt.Println("products[:3] => ", products[:3])
	fmt.Println("products[3:] => ", products[3:])

	//Creating a copy of slice (aka underlying array)
	copyOfProducts := make([]string, len(products), 10)
	copy(copyOfProducts, products)
	fmt.Println("After copying")
	fmt.Println(products, copyOfProducts)
	products[0] = "FountainPen"
	fmt.Println(products, copyOfProducts)

	copyOfProducts = append(copyOfProducts, "BallPointPen")
	fmt.Println(products, copyOfProducts)

	products = append(products, "Paper")
	fmt.Println(products, copyOfProducts)

	//Maps
	fmt.Println("")
	fmt.Println("Maps")
	//var productRanks map[string]int = map[string]int{}
	productRanks := map[string]int{
		"Pen":    1,
		"Pencil": 2,
		"Marker": 5,
		"Eraser": 3,
	}
	fmt.Println(productRanks)

	//iterating a map
	fmt.Println("Iterating a map")
	for key, value := range productRanks {
		fmt.Println(key, value)
	}

	//adding a new key-value pair
	productRanks["ScribblePad"] = 4
	fmt.Println("After adding a new key-value pair")
	fmt.Println(productRanks)

	//verifying if a key exists
	if rank, ok := productRanks["Printer"]; ok {
		fmt.Println("Printer is ranked with", rank)
	} else {
		fmt.Println("Printer is not ranked")
	}

	rankOfPrinter := productRanks["Printer"]
	fmt.Println("Rank of Printer => ", rankOfPrinter)

	//Deleting a key-value pair
	delete(productRanks, "Eraser")
	fmt.Println("After deleting a key-value pair")
	fmt.Println(productRanks)
}
