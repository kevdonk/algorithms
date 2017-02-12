package main

import "fmt"

func rotate(nums []int) []int {
	last := len(nums) - 1
	numberToRotate := nums[last]
	for i := (last); i > 0; i-- {
		nums[i] = nums[i-1]
	}
	nums[0] = numberToRotate
	return nums
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
	for i := 0; i < rotations; i++ {
		numbers = rotate(numbers)
	}
	for i := 0; i < queries; i++ {
		fmt.Scanf("%v", &m)
		fmt.Println(query(numbers, m))
	}
}
