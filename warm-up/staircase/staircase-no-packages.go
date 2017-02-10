package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for j := 1; j < n+1; j++ {
			if n-j > i {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Print("\n")
	}
}
