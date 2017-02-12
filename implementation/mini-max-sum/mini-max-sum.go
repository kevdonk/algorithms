package main

import (
	"fmt"
	"math"
)

func main() {
	var numbers [5]int64
	var sums [5]int64
	var min, max int64
	fmt.Scanf("%v %v %v %v %v", &numbers[0], &numbers[1], &numbers[2], &numbers[3], &numbers[4])
	for i := range numbers {
		for x := 0; x < 4; x++ {
			sums[i] += numbers[int64(math.Mod(float64(i+x), 5))]
		}
		if min == 0 {
			min = sums[i]
		}
		if sums[i] < min {
			min = sums[i]
		}
		if sums[i] > max {
			max = sums[i]
		}
	}
	fmt.Printf("%v %v", min, max)
}
