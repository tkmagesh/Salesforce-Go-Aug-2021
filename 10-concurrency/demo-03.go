package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex = sync.Mutex{}

//communicate by sharing memory [NOT ADVICABLE in GOLANG]

var operCount = 0

func main() {
	var wg *sync.WaitGroup = &sync.WaitGroup{}
	wg.Add(1)
	go print("Hello", wg)

	wg.Add(1)
	go print("World", wg)

	wg.Wait()
	fmt.Println(operCount)
	fmt.Println("Exiting from main")
}

func print(str string, wg *sync.WaitGroup) {
	fmt.Println(str)
	mutex.Lock()
	{
		operCount++
	}
	mutex.Unlock()
	wg.Done()
}
