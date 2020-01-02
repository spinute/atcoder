package main

import "fmt"

func sum_digits(n int) int {
	if n < 10 {
		return n
	} else {
		return n%10 + sum_digits(n/10)
	}
}

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	ans := 0
	for i := 1; i <= n; i++ {
		if s := sum_digits(i); a <= s && s <= b {
			ans += i
		}
	}
	fmt.Println(ans)
}
