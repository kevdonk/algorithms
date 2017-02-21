package main

import "fmt"

func between(x, left, right int) bool {
	return x >= left && x <= right
}

func main() {
	var houseLeft, houseRight, appleTree, orangeTree, appleCount, orangeCount, distance, applesHit, orangesHit int
	var appleHitMin, appleHitMax, orangeHitMin, orangeHitMax int
	fmt.Scanln(&houseLeft, &houseRight)
	fmt.Scanln(&appleTree, &orangeTree)
	fmt.Scanln(&appleCount, &orangeCount)
	appleHitMin = houseLeft - appleTree
	appleHitMax = houseRight - appleTree
	orangeHitMin = houseLeft - orangeTree
	orangeHitMax = houseRight - orangeTree
	for i := 0; i < appleCount; i++ {
		fmt.Scan(&distance)
		if between(distance, appleHitMin, appleHitMax) {
			applesHit++
		}
	}
	for i := 0; i < orangeCount; i++ {
		fmt.Scan(&distance)
		if between(distance, orangeHitMin, orangeHitMax) {
			orangesHit++
		}
	}
	fmt.Println(applesHit)
	fmt.Println(orangesHit)
}
