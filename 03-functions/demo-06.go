package main

import "fmt"

func main() {

	increment, decrement := getCounter()
	fmt.Println("Incrementing")
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println("Decrementing")
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())

}

func getCounter() (func() int, func() int) { //step - 1
	var count int = 0         //step - 2
	increment := func() int { //step - 3
		count += 1 //step - 4
		return count
	}
	decrement := func() int { //step - 3
		count -= 1 //step - 4
		return count
	}
	return increment, decrement
}
