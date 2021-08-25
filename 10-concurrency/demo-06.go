package main

import (
	"fmt"
)

//Share memory by communicating

func main() {

	ch := make(chan int)
	opCountCh := make(chan int)

	go add(100, 200, ch, opCountCh)
	opCountCh <- 0
	//reading data from the channel

	count := <-opCountCh
	result := <-ch
	fmt.Println(result, count)
}

func add(x, y int, ch chan int, opCountCh chan int) {
	count := <-opCountCh
	result := x + y

	//writing data into the channel
	//the following line can succeed only when there is read operation initiated and pending
	count++
	opCountCh <- count
	ch <- result

}
