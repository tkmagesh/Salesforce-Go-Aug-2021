package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go fibonacci(c)
	time.Sleep(2 * time.Second)
	for i := range c {
		println(i)
	}
	fmt.Println("Done")
}

func fibonacci(ch chan int) {
	x, y := 1, 1
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
