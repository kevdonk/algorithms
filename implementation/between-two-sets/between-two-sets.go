package main

import (
	"fmt"
	"sort"
)

func isFactor(small, big int) bool {
	return big%small == 0
}

func getLowestCommonMultiple(series []int, size int) int {
	var lastIndex, lcm int
	sort.Ints(series)
	lastIndex = size - 1
	lcm = series[lastIndex]

	for i := lastIndex; i > 1; i-- {
		if !isFactor(series[i-1], lcm) {
			lcm = lcm * series[lastIndex-1] / getGreatestCommonDenominator(lcm, series[lastIndex-1])
		}
	}

	return lcm
}

func getGreatestCommonDenominator(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func main() {
	var aSize, bSize, lcm, count, bLast int
	fmt.Scanf("%v %v", &aSize, &bSize)
	var a = make([]int, aSize)
	var b = make([]int, bSize)
	for i := 0; i < aSize; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < bSize; i++ {
		fmt.Scan(&b[i])
	}

	lcm = getLowestCommonMultiple(a, aSize)
	bLast = bSize - 1

	for i := lcm; i <= b[0]; i += lcm {
		for j := 0; j <= bLast; j++ {
			if isFactor(i, b[j]) {
				if j == bLast {
					count++
				}
			} else {
				j = bLast
			}
		}
	}

	fmt.Print(count)
}
