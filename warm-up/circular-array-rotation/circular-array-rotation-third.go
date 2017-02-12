package main

import "fmt"

func main() {
	var n, rotations, queries, m int
	fmt.Scan(&n, &rotations, &queries)
	var numbers = make([]int, n)
	for i := range numbers {
		fmt.Scan(&numbers[(i+rotations)%n])
	}
	for i := 0; i < queries; i++ {
		fmt.Scan(&m)
		fmt.Println(numbers[m])
	}
}
