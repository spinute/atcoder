package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	var ps, ys []int
	for i := 0; i < m; i++ {
		var p, y int
		fmt.Scan(&p, &y)
		ps = append(ps, p)
		ys = append(ys, y)
	}

	//p2cs := make(map[int][]int)
	var p2cs [100001][]int
	for i := 0; i < m; i++ {
		ar := p2cs[ps[i]]
		p2cs[ps[i]] = append(ar, ys[i])
	}

	// sorted_m := make(map[int][]int)
	var sorted_m [100001][]int
	for k, v := range p2cs {
		sort.Ints(v)
		sorted_m[k] = v
	}

	for i := 0; i < m; i++ {
		p := ps[i]
		y := ys[i]
		r := rank(sorted_m[p], y)
		format(p, r)
	}
}

func format(p, r int) {
	fmt.Printf("%06d%06d\n", p, r)
}

func rank(sorted []int, v int) int {
	l := 0
	r := len(sorted)
	for l+1 < r {
		m := (l + r) / 2
		if v < sorted[m] {
			r = m
		} else {
			l = m
		}
	}
	return l + 1
}
