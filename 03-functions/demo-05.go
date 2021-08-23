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

	wrappedAdd := wrappedInvoke(add)
	wrappedSubtract := wrappedInvoke(subtract)

	wrappedAdd(100, 200)
	wrappedSubtract(100, 200)
}

func wrappedInvoke(oper func(int, int) int) func(int, int) {
	return func(x int, y int) {
		fmt.Println("Before invocation")
		fmt.Println(oper(x, y))
		fmt.Println("After invocation")
	}
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}
