package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	go fn(ch)
	fmt.Println("[@main] attempting to read 10")
	fmt.Println(<-ch)
	fmt.Println("[@main] finished reading 10")
	fmt.Println("[@main] attempting to read 20")
	fmt.Println(<-ch)
	fmt.Println("[@main] finished reading 20")
	fmt.Println("[@main] attempting to read 30")
	fmt.Println(<-ch)
	fmt.Println("[@main] finished reading 30")
	fmt.Println("[@main] attempting to read 40")
	fmt.Println(<-ch)
	fmt.Println("[@main] finished reading 40")
	fmt.Println("[@main] attempting to read 50")
	fmt.Println(<-ch)
	fmt.Println("[@main] finished reading 50")
	fmt.Println("[@main] attempting to read 60")
	fmt.Println(<-ch)
	fmt.Println("[@main] finished reading 60")
	fmt.Println("[@main] attempting to read 70")
	fmt.Println(<-ch)
	fmt.Println("[@main] finished reading 70")
}

func fn(ch chan int) {
	fmt.Println("   [@fn] attempting to write 10")
	ch <- 10
	fmt.Println("   [@fn] finished writing 10")
	fmt.Println("   [@fn] attempting to write 20")
	ch <- 20
	fmt.Println("   [@fn] finished writing 20")
	fmt.Println("   [@fn] attempting to write 30")
	ch <- 30
	fmt.Println("   [@fn] finished writing 30")
	fmt.Println("   [@fn] attempting to write 40")
	ch <- 40
	fmt.Println("   [@fn] finished writing 40")
	fmt.Println("   [@fn] attempting to write 50")
	ch <- 50
	fmt.Println("   [@fn] finished writing 50")
	fmt.Println("   [@fn] attempting to write 60")
	ch <- 60
	fmt.Println("   [@fn] finished writing 60")
	fmt.Println("   [@fn] attempting to write 70")
	ch <- 70
	fmt.Println("   [@fn] finished writing 70")
}
