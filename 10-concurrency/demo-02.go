package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg *sync.WaitGroup = &sync.WaitGroup{}
	wg.Add(1)
	go print("Hello", wg)

	wg.Add(1)
	go print("World", wg)
	wg.Wait()

	fmt.Println("Exiting from main")
}

func print(str string, wg *sync.WaitGroup) {
	fmt.Println(str)
	wg.Done()
}
