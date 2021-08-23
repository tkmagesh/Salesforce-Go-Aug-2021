package main

import "fmt"

func main() {
	//if construct
	/*
		no := 21
		if no%2 == 0 {
			fmt.Println(no, "is an even number")
		} else {
			fmt.Println(no, "is an odd number")
		}
	*/

	// no variable is scoped to the if block (i.e., not accessible outside of if block)
	if no := 21; no%2 == 0 {
		fmt.Println(no, "is an even number")
	} else {
		fmt.Println(no, "is an odd number")
	}

	//for construct
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//simulating a 'while' construct using 'for'
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Println(numSum)

	//infinite loop
	x := 1
	for {
		x += x
		if x > 100 {
			break
		}
	}
	fmt.Println(x)

	//switch construct
	/*
		no := 20
		switch no % 2 {
		case 0:
			fmt.Println(no, "is an even number")
		case 1:
			fmt.Println(no, "is an odd number")
		}
	*/

	/*
		switch no := 20; no % 2 {
		case 0:
			{
				fmt.Println(no, "is an even number")
			}
		case 1:
			{
				fmt.Println(no, "is an odd number")
			}
		}
	*/
	no := 20
	switch {
	case no%2 == 0:
		fmt.Println(no, "is an even number")
	case no%2 == 1:
		fmt.Println(no, "is an odd number")
	}

	/*
		rank by score
		score 0 - 3 => "Terrible"
		score 4 - 7 => "Not bad"
		score 8 - 9 => "Good"
		score 10 => "Excellent"
		otherwise => "Invalid score"
	*/

	/*
		var score int
		fmt.Println("Enter the score :")
		fmt.Scanf("%d", &score)
	*/
	var score int = 8
	switch score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not bad")
	case 8, 9:
		fmt.Println("Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid score")
	}

	//using fallthrough
	n := 7
	switch n {
	case 0:
		fmt.Println("n is 0")
		fallthrough
	case 1:
		fmt.Println("n <= 1")
		fallthrough
	case 2:
		fmt.Println("n <= 2")
		fallthrough
	case 3:
		fmt.Println("n <= 3")
		fallthrough
	case 4:
		fmt.Println("n <= 4")
		fallthrough
	case 5:
		fmt.Println("n <= 5")
		fallthrough
	case 6:
		fmt.Println("n <= 6")
		//fallthrough
	case 7:
		fmt.Println("n <= 7")
		fallthrough
	case 8:
		fmt.Println("n <= 8")

	}
}
