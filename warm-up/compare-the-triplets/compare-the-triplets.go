package main

import "fmt"

type Score struct {
	Clarity     int
	Originality int
	Difficulty  int
}

func compareScore(a Score, b Score) int {
	var score int
	if a.Clarity > b.Clarity {
		score += 1
	}
	if a.Originality > b.Originality {
		score += 1
	}
	if a.Difficulty > b.Difficulty {
		score += 1
	}
	return score
}

func main() {
	var a, b Score
	fmt.Scanf("%d %d %d", &a.Clarity, &a.Originality, &a.Difficulty)
	fmt.Scanf("%d %d %d", &b.Clarity, &b.Originality, &b.Difficulty)
	fmt.Println(compareScore(a, b), compareScore(b, a))
}
