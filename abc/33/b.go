package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	ss := make([]string, n)
	ps := make([]int, n)
	for i := range ps {
		fmt.Scan(&ss[i], &ps[i])
	}

	sum := 0
	for _, v := range ps {
		sum += v
	}

	ans := ""
	for i, v := range ps {
		if v > sum/2 {
			ans = ss[i]
		}
	}

	if ans == "" {
		ans = "atcoder"
	}

	fmt.Println(ans)
}
