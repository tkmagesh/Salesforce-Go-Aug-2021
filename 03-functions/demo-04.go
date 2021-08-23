/* functions are arguments */

package main

import "fmt"

func main() {
	/*
		fmt.Println("Before invocation")
		fmt.Println(add(100, 200))
		fmt.Println("After invocation")

		fmt.Println("Before invocation")
		fmt.Println(subtract(100, 200))
		fmt.Println("After invocation")
	*/

	wrapInvoke(add, 100, 200)
	wrapInvoke(subtract, 100, 200)
}

func wrapInvoke(oper func(int, int) int, x, y int) {
	fmt.Println("Before invocation")
	fmt.Println(oper(x, y))
	fmt.Println("After invocation")
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}
