package main

import "fmt"

func main() {
	var V, n int

	fmt.Scan(&V)
	fmt.Scan(&n)
	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if a[i] == V {
			fmt.Println(i)
		}
	}
}
