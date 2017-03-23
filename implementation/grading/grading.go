package main

import "fmt"

func round(grade int) int {
	if grade > 37 {
		remainder := grade % 5
		if remainder > 2 {
			return grade + (5 - remainder)
		}
	}
	return grade
}

func main() {
	var students, grade int
	fmt.Scan(&students)
	for i := 0; i < students; i++ {
		fmt.Scan(&grade)
		fmt.Println(round(grade))
	}
}
