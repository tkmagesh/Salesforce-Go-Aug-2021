/* panic & recovery */
package main

import (
	"errors"
	"fmt"
)

var divideByZero = errors.New("division by zero")

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main:", r)
			fmt.Println("Exiting from main")
		}
	}()
	fmt.Println(divide(10, 2))
	fmt.Println(divide(10, 0))

}

func divide(x, y int) (result int, err error) {
	defer func() {
		if e := recover(); e != nil {
			if e == divideByZero {
				fmt.Println("Recovered from 'division by zero'")
			} else {
				fmt.Println("Recovered in divide:", e)
				err = errors.New("Divide operation failed")
			}
		}
	}()
	if y == 0 {
		panic(divideByZero)
	}
	result = x / y
	return
}
