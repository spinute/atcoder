package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n, &q)
	ls := make([]int, q)
	rs := make([]int, q)
	ts := make([]int, q)
	for i := range ls {
		fmt.Scan(&ls[i], &rs[i], &ts[i])
	}

	as := make([]int, n)
	for i := range ls {
		for j := ls[i]; j <= rs[i]; j++ {
			as[j-1] = ts[i]
		}
	}

	for _, a := range as {
		fmt.Println(a)
	}
}
