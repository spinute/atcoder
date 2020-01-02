package main

import "fmt"

func check(b byte) string {
	switch b {
	case 'a', 'e', 'i', 'o', 'u':
		return "vowel"
	}
	return "consonant"
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(check(s[0]))
}
