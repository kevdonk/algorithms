package main

import (
	"fmt"
)

func main() {
	var min, max, sum, number int64
	for i := 0; i < 5; i++ {
		fmt.Scan(&number)
		if min == 0 {
			min = number
		}
		if number < min {
			min = number
		}
		if number > max {
			max = number
		}
		sum += number
	}
	fmt.Println(sum-max, sum-min)
}
