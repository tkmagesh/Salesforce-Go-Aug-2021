package main

import "fmt"

func main() {
	fmt.Println(sum())                                //=> 0
	fmt.Println(sum(10))                              //=> 10
	fmt.Println(sum(10, 20))                          //=> 30
	fmt.Println(sum(10, 20, 30, 40))                  //=> 100
	fmt.Println(sum(10, 20, "30", 40))                //=> 100
	fmt.Println(sum(10, 20, "abc", 40))               //=> 70
	fmt.Println(sum(10, 20, []int{30, 40}))           //=> 100
	fmt.Println(sum(10, 20, []interface{}{"30", 40})) //=> 100
}

func sum( /*  */ ) int {
	return 0
}
