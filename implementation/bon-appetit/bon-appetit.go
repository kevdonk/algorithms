package main

import "fmt"

func main() {
	var n, k, paid, sum int
	fmt.Scan(&n, &k)
	costs := make([]int, n)
	for i := range costs {
		fmt.Scan(&costs[i])
		if i != k {
			sum += costs[i]
		}
	}
	fmt.Scan(&paid)
	owed := sum / 2
	if (owed) == paid {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(paid - owed)
	}
}
