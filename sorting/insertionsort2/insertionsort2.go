package main

import "fmt"

func printArray(array []int) {
	for j := range array {
		fmt.Printf("%v ", array[j])
	}
	fmt.Println()
}

func sortArray(array []int, start int) []int {
	var i int
	toInsert := array[start]
	for i = start - 1; i >= 0; i-- {
		if array[i] > toInsert {
			array[i+1] = array[i]
		} else {
			break
		}
	}
	array[i+1] = toInsert
	return array
}

func main() {
	var size, i int
	fmt.Scan(&size)
	array := make([]int, size)
	for i = 0; i < size; i++ {
		fmt.Scan(&array[i])
	}
	for i = 1; i < size; i++ {
		array = sortArray(array, i)
		printArray(array)
	}
}
