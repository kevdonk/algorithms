package main

import "fmt"

func main() {
	var size, sum, number int
	fmt.Scanf("%v\n", &size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%v", &number)
		sum += number
	}
	fmt.Println(sum)
}
