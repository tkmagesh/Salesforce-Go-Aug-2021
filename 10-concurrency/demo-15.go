package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	stop := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			fmt.Print("tick")
		case <-stop:
			fmt.Println("Done")
			return
		default:
			fmt.Print(".")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
