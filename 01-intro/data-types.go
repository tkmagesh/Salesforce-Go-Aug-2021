package main

import "fmt"

/* := syntax not applicable here */
var msg = "Hello, world!"

func main() {

	/*
		var msg string
		msg = "Hello World!"
	*/

	/*
		var msg string = "Hello World!"
	*/

	/*
		// using the type inference
		var msg = "Hello World!"
	*/

	//the following is applicable only in a function scoped variable (NOT int package scope)
	//msg := "Hello World!"

	fmt.Println(msg)

	/* Declaring multiple variables of different types */

	/*
		var x int
		var y int
		var str string
		x = 100
		y = 200
		str = "Add Result = "
	*/

	/*
		var x, y int
		var str string
	*/
	/*
		var (
			x, y int
			str  string
		)
		x = 100
		y = 200
		str = "Add Result = "
	*/

	/*
		var x int = 100
		var y int = 200
		var str string = "Add Result = "
	*/

	/*
		var (
			x   int    = 100
			y   int    = 200
			str string = "Add Result = "
		)
	*/
	/*
		var (
			x, y int    = 100, 200
			str  string = "Add Result = "
		)
	*/
	x, y, str := 100, 200, "Add Result ="

	fmt.Println(str, x+y)

	var t bool = true
	fmt.Println(t)

	var noInt int = 100
	//type conversion
	var noFloat float64 = float64(noInt)
	fmt.Println(noInt, noFloat)
}
