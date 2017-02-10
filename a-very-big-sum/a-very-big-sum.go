package main

import "fmt"

func main() {
	var numberOfInts int
	var sum int64
	fmt.Scanf("%v\n", &numberOfInts)
	for i := 0; i < numberOfInts; i++ {
		var n int64
		fmt.Scanf("%v", &n)
		sum += n
	}
	fmt.Printf("%v\n", sum)
}
