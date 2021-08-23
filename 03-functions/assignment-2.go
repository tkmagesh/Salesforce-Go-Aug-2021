package main

import "fmt"

func main() {
	getCounter()()
	counter := getCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	anotherCounter := getCounter()
	fmt.Println(anotherCounter())
}

func getCounter() func() int { //step - 1
	var count int = 0       //step - 2
	counter := func() int { //step - 3
		count += 1 //step - 4
		return count
	}
	return counter //step - 5
}
