package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	as := make([]int, n)
	for i := range as {
		fmt.Scan(&as[i])
	}

	ans := 0
	if v := as[0] - x; v > 0 {
		as[0] -= v
		ans += v
	}
	for i := 0; i < len(as)-1; i++ {
		if v := as[i] + as[i+1] - x; v > 0 {
			ans += v
			as[i+1] -= v
		}
	}
	fmt.Println(ans)
}
