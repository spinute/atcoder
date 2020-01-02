package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Scan(&n)
	ts := make([]int, n)
	as := make([]int, n)
	for i := range as {
		fmt.Scan(&ts[i], &as[i])
	}

	a := as[0]
	t := ts[0]
	for i := range ts {
		d := max((t-1)/ts[i]+1, (a-1)/as[i]+1)
		a = d * as[i]
		t = d * ts[i]
	}
	fmt.Println(a + t)
}
