package main

import "fmt"

func main() {
	var n, primary, secondary, sum int
	fmt.Scanf("%v", &n)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for ii := range matrix[i] {
			fmt.Scanf("%v", &matrix[i][ii])
			if i == ii {
				primary += matrix[i][ii]
			}
			if i+ii == (n - 1) {
				secondary += matrix[i][ii]
			}
		}
	}
	sum = primary - secondary
	if sum < 0 {
		sum = -sum
	}
	fmt.Printf("%v", sum)
}
