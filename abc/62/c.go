package main

import "fmt"

func max(vs []int) int {
	m := vs[0]
	for _, v := range vs {
		if v > m {
			m = v
		}
	}
	return m
}
func min(vs []int) int {
	m := vs[0]
	for _, v := range vs {
		if v < m {
			m = v
		}
	}
	return m
}
func score(a, b, c int) int {
	return max([]int{a, b, c}) - min([]int{a, b, c})
}

func c1(a, b int) int {
	v1 := a / 3 * b
	v2 := (a - a/3) / 2 * b
	v3 := a*b - v1 - v2
	return score(v1, v2, v3)
}
func c2(a, b int) int {
	v1 := a / 3 * b
	a_rem := a - a/3
	v2 := b / 2 * a_rem
	v3 := a*b - v1 - v2
	return score(v1, v2, v3)
}
func c3(a, b int) int {
	v1 := (a/3 + 1) * b
	a_rem := a - a/3 - 1
	v2 := b / 2 * a_rem
	v3 := a*b - v1 - v2
	return score(v1, v2, v3)
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	fmt.Println(min([]int{c1(h, w), c1(w, h), c2(h, w), c2(w, h), c3(h, w), c3(w, h)}))
}
