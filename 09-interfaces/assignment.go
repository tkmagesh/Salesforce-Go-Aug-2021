package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum())                                               //=> 0
	fmt.Println(sum(10))                                             //=> 10
	fmt.Println(sum(10, 20))                                         //=> 30
	fmt.Println(sum(10, 20, 30, 40))                                 //=> 100
	fmt.Println(sum(10, 20, "30", 40))                               //=> 100
	fmt.Println(sum(10, 20, "abc", 40))                              //=> 70
	fmt.Println(sum(10, 20, []int{30, 40}))                          //=> 100
	fmt.Println(sum(10, 20, []interface{}{"30", 40}))                //=> 100
	fmt.Println(sum(10, 20, []interface{}{"30", 40, []int{50, 60}})) //=> 210
}

func sum(values ...interface{}) int {
	result := 0
	for _, no := range values {
		switch value := no.(type) {
		case int:
			result += value
		case string:
			if val, err := strconv.Atoi(value); err == nil {
				result += val
			}
		case []interface{}:
			result += sum(value...)
		case []int:
			intfList := make([]interface{}, len(value))
			for i, v := range value {
				intfList[i] = v
			}
			result += sum(intfList...)
		}
	}
	return result
}
