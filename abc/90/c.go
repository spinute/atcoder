package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	ans := 0
	switch {
	case n == 1 && m == 1:
		ans = 1
	case n == 1:
		ans = m - 2
	case m == 1:
		ans = n - 2
	default:
		ans = (n - 2) * (m - 2)
	}
	fmt.Println(ans)
}
