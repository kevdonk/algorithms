// http://rosettacode.org/wiki/Reverse_a_string#Go

// Interesting reading on reversing a string in Go and dealing with UTF-8 combining characters

package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// no encoding
func reverseBytes(s string) string {
	r := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		r[i] = s[len(s)-1-i]
	}
	return string(r)
}

// reverseCodePoints interprets its argument as UTF-8 and ignores bytes
// that do not form valid UTF-8.  return value is UTF-8.
func reverseCodePoints(s string) string {
	r := make([]rune, len(s))
	start := len(s)
	for _, c := range s {
		// quietly skip invalid UTF-8
		if c != utf8.RuneError {
			start--
			r[start] = c
		}
	}
	return string(r[start:])
}

// reversePreservingCombiningCharacters interprets its argument as UTF-8
// and ignores bytes that do not form valid UTF-8.  return value is UTF-8.
func reversePreservingCombiningCharacters(s string) string {
	if s == "" {
		return ""
	}
	p := []rune(s)
	r := make([]rune, len(p))
	start := len(r)
	for i := 0; i < len(p); {
		// quietly skip invalid UTF-8
		if p[i] == utf8.RuneError {
			i++
			continue
		}
		j := i + 1
		for j < len(p) && (unicode.Is(unicode.Mn, p[j]) ||
			unicode.Is(unicode.Me, p[j]) || unicode.Is(unicode.Mc, p[j])) {
			j++
		}
		for k := j - 1; k >= i; k-- {
			start--
			r[start] = p[k]
		}
		i = j
	}
	return (string(r[start:]))
}

func main() {
	test("asdf")
	test("as⃝df̅")
}

func test(s string) {
	fmt.Println("\noriginal:      ", []byte(s), s)
	r := reverseBytes(s)
	fmt.Println("reversed bytes:", []byte(r), r)
	fmt.Println("original code points:", []rune(s), s)
	r = reverseCodePoints(s)
	fmt.Println("reversed code points:", []rune(r), r)
	r = reversePreservingCombiningCharacters(s)
	fmt.Println("combining characters:", []rune(r), r)
}

// Output:
// original:       [97 115 100 102] asdf
// reversed bytes: [102 100 115 97] fdsa
// original code points: [97 115 100 102] asdf
// reversed code points: [102 100 115 97] fdsa
// combining characters: [102 100 115 97] fdsa

// original:       [97 115 226 131 157 100 102 204 133] as⃝df̅
// reversed bytes: [133 204 102 100 157 131 226 115 97] ��fd���sa
// original code points: [97 115 8413 100 102 773] as⃝df̅
// reversed code points: [773 102 100 8413 115 97] ̅fd⃝sa
// combining characters: [102 773 100 115 8413 97] f̅ds⃝a
