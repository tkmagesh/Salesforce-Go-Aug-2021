package main

import (
	"fmt"
)

//Share memory by communicating

type Data struct {
	result int
	count  int
}

func main() {

	ch := make(chan Data)
	data := Data{}

	go add(100, 200, ch)
	ch <- data
	//reading data from the channel

	r := <-ch
	fmt.Println(r.result, r.count)
}

func add(x, y int, ch chan Data) {
	d := <-ch
	d.result = x + y
	d.count = d.count + 1
	//writing data into the channel
	//the following line can succeed only when there is read operation initiated and pending

	ch <- d

}
