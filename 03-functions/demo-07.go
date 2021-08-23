//errors

package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("Invalid argument. Cannot divide by zero.")

func main() {
	fmt.Println(divide(10, 2))
	result, err := divide(10, 0)
	if err == DivideByZeroError {
		fmt.Println("Divide by zero error.")
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

}

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, DivideByZeroError
	}
	return x / y, nil

}
