package main

import "fmt"

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int64) int64 {
	return a / gcd(a, b) * b
}

func main() {
	var n int64
	fmt.Scan(&n)
	ts := make([]int64, n)
	for i := range ts {
		fmt.Scan(&ts[i])
	}

	ans := ts[0]
	for _, t := range ts {
		ans = lcm(ans, t)
	}
	fmt.Println(ans)
}
