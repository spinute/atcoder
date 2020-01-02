package main

import "fmt"

func solve(a, b, c int) bool {
	for i := 0; i < b; i++ {
		if a*i%b == c {
			return true
		}
	}
	return false
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if solve(a, b, c) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
