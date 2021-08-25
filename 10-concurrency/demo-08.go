package main

import (
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go print("Hello", 1*time.Second, ch1, ch2, wg)
	wg.Add(1)
	go print("World", 2*time.Second, ch2, ch1, wg)
	ch1 <- 1
	wg.Wait()

}

func print(msg string, delay time.Duration, ch1, ch2 chan int, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		<-ch1
		time.Sleep(delay)
		println(msg)
		ch2 <- 1
	}
	wg.Done()
}
