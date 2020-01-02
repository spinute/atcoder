package main

import "fmt"

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func dist(x1, y1, x2, y2 int) int {
	return diff(x1, x2) + diff(y1, y2)
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	as := make([]int, n)
	bs := make([]int, n)
	cs := make([]int, m)
	ds := make([]int, m)
	for i := range as {
		fmt.Scan(&as[i], &bs[i])
	}
	for i := range cs {
		fmt.Scan(&cs[i], &ds[i])
	}

	for i := range as {
		a := as[i]
		b := bs[i]
		min := dist(a, b, cs[0], ds[0])
		min_j := 0
		for j := range cs {
			c := cs[j]
			d := ds[j]
			if v := dist(a, b, c, d); v < min {
				min = v
				min_j = j
			}
		}
		fmt.Println(min_j + 1)
	}
}
