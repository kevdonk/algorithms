// see constraint x.1 <= x.2
package main

import "fmt"

func main() {
	var x1, x2, v1, v2 int
	fmt.Scanf("%v %v %v %v", &x1, &v1, &x2, &v2)
	if v2 < v1 && ((x2-x1)%(v1-v2) == 0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
