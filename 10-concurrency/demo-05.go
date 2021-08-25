package main

import (
	"fmt"
	"sync"
)

//Share memory by communicating

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ch := make(chan int)
	go add(100, 200, ch, wg)
	wg.Wait()

	//reading data from the channel
	result := <-ch
	fmt.Println(result)
}

func add(x, y int, ch chan int, wg *sync.WaitGroup) {
	result := x + y

	//writing data into the channel
	ch <- result
	wg.Done()
}
