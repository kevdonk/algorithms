package main

import "fmt"

func main() {
	var n, k, count int
	fmt.Scan(&n, &k)
	numbers := make([]int, n)
	for x := 0; x < n; x++ {
		fmt.Scan(&numbers[x])
	}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if (numbers[i]+numbers[j])%k == 0 {
				count++
			}
		}
	}
	fmt.Print(count)
}
