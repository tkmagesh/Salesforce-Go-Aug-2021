package main

import (
	"fmt"
	"sync"
)

type OpCount struct {
	count int
	sync.Mutex
}

func (opCount *OpCount) increment() {
	opCount.Lock()
	{
		opCount.count++
	}
	opCount.Unlock()
}

var opCount = &OpCount{}

func main() {
	var wg *sync.WaitGroup = &sync.WaitGroup{}

	wg.Add(1)
	go print("Hello", wg)

	wg.Add(1)
	go print("World", wg)

	wg.Wait()
	fmt.Println(opCount.count)
	fmt.Println("Exiting from main")
}

func print(str string, wg *sync.WaitGroup) {
	fmt.Println(str)
	opCount.increment()
	wg.Done()
}
