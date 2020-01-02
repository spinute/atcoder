package main

import "fmt"

func s(n int) int {
	if n == 0 {
		return 0
	} else {
		return n%10 + s(n/10)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	if n%s(n) == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
