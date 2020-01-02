package main

import "fmt"

func check(s, t string) bool {
	l := len(s)
	for i := range s {
		if s[i] != t[l-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	if check(s, t) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
