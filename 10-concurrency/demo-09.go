package main

import "fmt"

func main() {
	data := []int{4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55}
	dataSet1 := data[:len(data)/2]
	dataSet2 := data[len(data)/2:]
	resultCh := make(chan int)
	go sum(resultCh, dataSet1...)
	go sum(resultCh, dataSet2...)
	result := <-resultCh + <-resultCh
	fmt.Println(result)
}

func sum(ch chan int, nos ...int) {
	var sum int
	for _, n := range nos {
		sum += n
	}
	ch <- sum
}
