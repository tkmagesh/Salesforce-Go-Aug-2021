package main

import "fmt"

func main() {
	var no int = 10
	var noPtr *int = &no
	fmt.Println(noPtr)

	//dereferencing -> accessing the value using the pointer
	var ptrValue = *noPtr
	fmt.Println(ptrValue)

	// no == *(&no)

	fmt.Println("Before incrementing, no => ", no)
	increment(&no)
	fmt.Println("After incrementing, no => ", no)

	n1, n2 := 10, 20
	fmt.Println("Before swapping", n1, n2)
	swap(&n1, &n2)
	fmt.Println("After swapping", n1, n2)
}

func increment(x *int) {
	*x += 1
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
