package main

import "fmt"

func printArray(array []int) {
	for j := range array {
		fmt.Printf("%v ", array[j])
	}
	fmt.Println()
}

func main() {
	var size, toInsert, i int
	fmt.Scan(&size)
	array := make([]int, size)
	for i = 0; i < size; i++ {
		fmt.Scan(&array[i])
	}
	toInsert = array[size-1]
	for i = size - 2; i >= 0; i-- {
		if array[i] > toInsert {
			array[i+1] = array[i]
			printArray(array)
		} else {
			break
		}
	}
	array[i+1] = toInsert
	printArray(array)
}
