package main

import (
	"fmt"
)

//Share memory by communicating

func main() {

	ch := make(chan int)
	go add(100, 200, ch)

	//reading data from the channel
	result := <-ch
	fmt.Println(result)
}

func add(x, y int, ch chan int) {
	result := x + y

	//writing data into the channel
	//the following line can succeed only when there is read operation initiated and pending
	ch <- result

}
