package main

import (
	"fmt"
	"strings"
)

func main() {
	var x int
	fmt.Scan(&x)
	for i := 1; i < x+1; i++ {
		// more efficient to use copy https://golang.org/pkg/builtin/#copy or buffers
		// http://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go
		step := strings.Repeat(" ", x-i) + strings.Repeat("#", i)
		fmt.Println(step)
	}
}
