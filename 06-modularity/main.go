package main

import (
	"fmt"
	calc "modularity_demo/calculator"
	"modularity_demo/calculator/utils"

	"github.com/fatih/color"
)

func main() {
	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Println(calc.OperationCount)

	color.Red("Is 97 a prime number ? : %t\n", utils.IsPrime(97))

}
