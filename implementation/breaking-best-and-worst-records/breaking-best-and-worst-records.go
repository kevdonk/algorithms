package main

import "fmt"

func main() {
	var games, breakHigh, breakLow int
	fmt.Scan(&games)
	scores := make([]int, games)
	fmt.Scan(&scores[0])
	high := scores[0]
	low := scores[0]
	for i := 1; i < games; i++ {
		fmt.Scan(&scores[i])
		if scores[i] > high {
			high = scores[i]
			breakHigh++
		} else if scores[i] < low {
			low = scores[i]
			breakLow++
		}
	}
	fmt.Printf("%v %v", breakHigh, breakLow)
}
