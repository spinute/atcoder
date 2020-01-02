package main

import "fmt"

func str(n int) string {
	if n == 0 {
		return "b"
	}
	switch n % 3 {
	case 1:
		return "a" + str(n-1) + "c"
	case 2:
		return "c" + str(n-1) + "a"
	case 0:
		return "b" + str(n-1) + "b"
	}
	return ""
}

func l2n(l int) int {
	return (l - 1) / 2
}

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	if n := l2n(len(s)); str(n) == s {
		fmt.Println(n)
	} else {
		fmt.Println(-1)
	}
}
