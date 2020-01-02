package main

import "fmt"

func max(vs []int) int {
	v := vs[0]
	for _, u := range vs {
		if u > v {
			v = u
		}
	}
	return v
}
func min(vs []int) int {
	v := vs[0]
	for _, u := range vs {
		if u < v {
			v = u
		}
	}
	return v
}

func main() {
	var n, m, x, y int
	fmt.Scan(&n, &m, &x, &y)
	xs := make([]int, n)
	ys := make([]int, m)
	for i := range xs {
		fmt.Scan(&xs[i])
	}
	for i := range ys {
		fmt.Scan(&ys[i])
	}

	xs = append(xs, x)
	ys = append(ys, y)

	if max(xs) < min(ys) {
		fmt.Println("No War")
	} else {
		fmt.Println("War")
	}
}
