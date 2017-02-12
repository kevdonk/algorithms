package main

import "fmt"

func rotate(nums []int, rotations int) []int {
	length := len(nums)
	newNums := make([]int, length)
	for i := range nums {
		j := i + rotations
		if j > length-1 {
			j = j % length
		}
		newNums[j] = nums[i]
	}
	return newNums
}

func query(nums []int, m int) int {
	return nums[m]
}

func main() {
	var n, rotations, queries, m int
	fmt.Scanf("%v %v %v", &n, &rotations, &queries)
	var numbers = make([]int, n)
	for i := range numbers {
		fmt.Scan(&numbers[i])
	}

	numbers = rotate(numbers, rotations)

	for i := 0; i < queries; i++ {
		fmt.Scanf("%v", &m)
		fmt.Println(query(numbers, m))
	}
}
