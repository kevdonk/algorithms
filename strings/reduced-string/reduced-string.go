package main

import "fmt"

func removePair(word string) string {
	for i := 0; i < len(word)-1; i++ {
		if word[i] == word[i+1] {
			newWord := word[0:i]
			if i+2 <= len(word)-1 {
				newWord = newWord + word[i+2:]
			}
			return removePair(newWord)
		}
	}
	if word == "" {
		return "Empty String"
	}
	return word
}

func main() {
	var word string
	fmt.Scan(&word)

	fmt.Println(removePair(word))
}
