package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	done := make(chan bool)
	go fibonacci(c, done)
	go func() {
		for i := range c {
			println(i)
		}
	}()
	fmt.Println("Hit Enter to stop")
	var input string
	fmt.Scanln(&input)
	done <- true
	fmt.Println("Done")
}

func fibonacci(ch chan int, done chan bool) {
	x, y := 1, 1
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
