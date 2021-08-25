package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go fibonacci(c)
	for i := range c {
		println(i)
	}
	fmt.Println("Done")
}

func fibonacci(ch chan int) {
	done := make(chan bool)
	x, y := 1, 1

	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case ch <- x:
			time.Sleep(500 * time.Millisecond)
			x, y = y, x+y
		case <-done:
			fmt.Println("Stopped")
			close(ch)
			return
		}
	}
}
