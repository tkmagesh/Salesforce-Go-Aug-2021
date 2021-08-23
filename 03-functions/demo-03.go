/* higher order functions */
package main

import "fmt"

func main() {
	//functions assigned as value to variables
	var fn func()
	fn = func() {
		fmt.Println("fn is invoked")
	}
	fn()

	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	fmt.Println(add(100, 200))

	//anonymous functions
	func() {
		fmt.Println("Anonymous function invoked")
	}()

	func(x, y int) {
		fmt.Println(x * y)
	}(10, 20)
}
