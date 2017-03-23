package main

import "fmt"

func main() {
	var count, bird, highest int
	birds := make(map[int]int)
	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		fmt.Scan(&bird)
		if birds[bird] > 0 {
			birds[bird]++
			if birds[bird] > birds[highest] {
				highest = bird
			} else if birds[bird] == birds[highest] {
				if bird < highest {
					highest = bird
				}
			}
		} else {
			birds[bird] = 1
		}
	}
	fmt.Println(highest)
}
