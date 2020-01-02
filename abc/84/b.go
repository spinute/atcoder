package main

import (
	"fmt"
	"unicode"
)

func are_digits(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func main() {
	var a, b int
	var s string
	fmt.Scan(&a, &b)
	fmt.Scan(&s)

	if are_digits(s[:a]) && are_digits(s[a+1:]) && s[a] == '-' {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
