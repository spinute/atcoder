package main

import "fmt"

func main() {
	var s string
	var n int
	fmt.Scan(&s, &n)
	ls := make([]int, n)
	rs := make([]int, n)
	for i := range ls {
		fmt.Scan(&ls[i], &rs[i])
	}

	is := make([]int, len(s))
	for i := range is {
		is[i] = i
	}
	for i := range ls {
		for l, r := ls[i]-1, rs[i]-1; l < r; l, r = l+1, r-1 {
			is[l], is[r] = is[r], is[l]
		}
	}

	ans := ""
	for _, i := range is {
		ans += string(s[i])
	}

	fmt.Println(ans)
}
