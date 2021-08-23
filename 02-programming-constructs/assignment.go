package main

import "fmt"

func main() {
	var userChoice, n1, n2, result int
	for {
		fmt.Println("Menu")
		fmt.Println("================================================================")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Println("Enter the choice :")
		fmt.Scanf("%d", &userChoice)
		switch userChoice {
		case 1, 2, 3, 4:
			fmt.Println("Enter two numbers :")
			fmt.Scanf("%d %d", &n1, &n2)
			switch userChoice {
			case 1:
				result = n1 + n2
			case 2:
				result = n1 - n2
			case 3:
				result = n1 * n2
			case 4:
				result = n1 / n2
			}
			fmt.Println("Result = ", result)
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
