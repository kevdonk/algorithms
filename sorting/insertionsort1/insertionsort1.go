package main

import "fmt"

func printArray(array []int) {
	for j := range array {
		fmt.Printf("%v ", array[j])
	}
	fmt.Println()
}

func main() {
	var size, toInsert int
	fmt.Scan(&size)
	array := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&array[i])
	}
	toInsert = array[size-1]
	for i := size - 2; i >= 0; i-- {
		if array[i] > toInsert {
			array[i+1] = array[i]
			if i == 0 {
				printArray(array)
				array[i] = toInsert
			}

		} else if array[i] < toInsert {
			array[i+1] = toInsert
			i = -1
		}
		printArray(array)
	}
}
