package main

import "fmt"

func main() {
	var n, q int
	var s string
	fmt.Scan(&n, &q, &s)
	ls := make([]int, q)
	rs := make([]int, q)
	for i := range ls {
		fmt.Scan(&ls[i], &rs[i])
	}

	ns := make([]int, n)
	for i := 1; i < n; i++ {
		if s[i-1:i+1] == "AC" {
			ns[i] = ns[i-1] + 1
		} else {
			ns[i] = ns[i-1]
		}
	}
	ms := make([]int, n)
	for i := 0; i < n-1; i++ {
		if s[i:i+2] == "AC" {
			if i == 0 {
				ms[i] = 1
			} else {
				ms[i] = ms[i-1] + 1
			}
		} else {
			if i == 0 {
				ms[i] = 0
			} else {
				ms[i] = ms[i-1]
			}
		}
	}

	for i, l := range ls {
		r := rs[i]
		if l == 1 {
			fmt.Println(ns[r-1])
		} else {
			fmt.Println(ns[r-1] - ms[l-2])
		}
	}
}
