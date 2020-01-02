package main

import "fmt"

func main() {
	var w, h, n int
	fmt.Scan(&w, &h, &n)
	xs := make([]int, n)
	ys := make([]int, n)
	as := make([]int, n)
	for i := range xs {
		fmt.Scan(&xs[i])
		fmt.Scan(&ys[i])
		fmt.Scan(&as[i])
	}

	xl := 0
	yl := 0
	xu := w
	yu := h
	for i, a := range as {
		x := xs[i]
		y := ys[i]
		switch a {
		case 1:
			if xl < x {
				xl = x
			}
		case 2:
			if xu > x {
				xu = x
			}
		case 3:
			if yl < y {
				yl = y
			}
		case 4:
			if yu > y {
				yu = y
			}
		}
	}
	if xl >= xu || yl >= yu {
		fmt.Println(0)
	} else {
		fmt.Println((xu - xl) * (yu - yl))
	}
}
