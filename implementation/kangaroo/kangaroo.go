package main

import "fmt"

func canBeCaught(ahead, aheadSpeed, behind, behindSpeed int) string {
	if aheadSpeed > behindSpeed {
		return "NO"
	} else if (ahead-behind)%(aheadSpeed-behindSpeed) == 0 {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	var x1, x2, v1, v2 int
	fmt.Scanf("%v %v %v %v", &x1, &v1, &x2, &v2)
	if x2 == x1 {
		fmt.Println("YES")
	}
	if (x2 != x1) && (v2 == v1) {
		fmt.Println("NO")
	}
	if x2 > x1 {
		fmt.Println(canBeCaught(x2, v2, x1, v1))
	} else {
		fmt.Println(canBeCaught(x1, v1, x2, v2))
	}
}
