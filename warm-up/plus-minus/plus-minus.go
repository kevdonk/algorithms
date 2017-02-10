package main

import "fmt"

func main() {
	var number, count int
	var positive, negative, zero float32
	fmt.Scanf("%v", &count)
	for i := 0; i < count; i++ {
		fmt.Scanf("%v", &number)
		if number < 0 {
			negative += 1
		}
		if number > 0 {
			positive += 1
		}
		if number == 0 {
			zero += 1
		}
	}
	count2 := float32(count)
	fmt.Printf("%v\n%v\n%v\n", positive/count2, negative/count2, zero/count2)
}
