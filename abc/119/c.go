package main

import "fmt"

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func score(mask int, ls []int, goal int) int {
	vs := make([]int, 0)
	for i := uint(0); 1<<i <= mask; i++ {
		if mask&(1<<i) > 0 {
			vs = append(vs, ls[i])
		}
	}

	min := 1000000
	mask2 := 1<<uint(len(vs)) - 1
	for i := uint(1); i <= uint(mask2); i++ {
		mp := 0
		sum := 0
		for j := uint(0); 1<<j <= mask2; j++ {
			if i&(1<<j) > 0 {
				sum += vs[j]
				mp += 10
			}
		}
		mp += diff(sum, goal)
		mp -= 10
		if mp < min {
			min = mp
		}
	}
	return min
}

func main() {
	var n, a, b, c int
	fmt.Scan(&n, &a, &b, &c)
	ls := make([]int, n)
	for i := range ls {
		fmt.Scan(&ls[i])
	}

	min := 1000000

	for i := 0; i < 1<<uint(n); i++ {
		for j := 0; j < 1<<uint(n); j++ {
			if i&j != 0 {
				continue
			}
			k := (1<<uint(n) - 1) &^ i &^ j

			if v := score(i, ls, a) + score(j, ls, b) + score(k, ls, c); v < min {
				min = v
			}
		}
	}

	fmt.Println(min)
}
