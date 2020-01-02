package main

import "fmt"

func same(a, b, c byte) bool {
	return a == b && b == c
}

func main() {
	var s string
	fmt.Scan(&s)
	if same(s[0], s[1], s[2]) || same(s[1], s[2], s[3]) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
