//variadic functions

package main

import "fmt"

func main() {
	fmt.Println(sum())
	fmt.Println(sum(10))
	fmt.Println(sum(10, 20))
	fmt.Println(sum(10, 20, 30, 40, 50))
}

func sum(nos ...int) int {
	/* nos => slice (array) of int */
	var result int
	/*
		for i := 0; i < len(nos); i++ {
			result += nos[i]
		}
	*/
	for _, no := range nos {
		result += no
	}
	return result
}
