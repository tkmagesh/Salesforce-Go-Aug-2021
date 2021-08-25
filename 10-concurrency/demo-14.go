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

	x, y := 1, 1
	/* done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}() */
	done := time.After(10 * time.Second)
	for {
		select {
		case ch <- x:
			time.Sleep(500 * time.Millisecond)
			x, y = y, x+y
		case t := <-done:
			fmt.Println(t)
			fmt.Println("Stopped")
			close(ch)
			return
		}
	}
}
