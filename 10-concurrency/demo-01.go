package main

import (
	"fmt"
)

func main() {
	go print("Hello")
	print("World")

	//time.Sleep(500 * time.Millisecond)
	//runtime.Gosched()
	fmt.Println("Exiting from main")
	var input string
	fmt.Scanln(&input)

}

func print(str string) {
	fmt.Println(str)
}
