package main

import "fmt"

func check(s string) bool {
	if s == "" {
		return true
	}
	if s[0] == 'o' || s[0] == 'k' || s[0] == 'u' {
		return check(s[1:])
	}
	if s[0:2] == "ch" {
		return check(s[2:])
	}
	return false
}

func main() {
	var s string
	fmt.Scan(&s)
	if check(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
