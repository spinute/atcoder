package main

import "fmt"

func main() {
	var n, a, b, k int
	fmt.Scan(&n, &a, &b, &k)
	ps := make([]int, k)
	for i := range ps {
		fmt.Scan(&ps[i])
	}

	counter := make(map[int]int)
	counter[a]++
	counter[b]++
	for _, p := range ps {
		counter[p]++
	}

	ans := "YES"
	for _, v := range counter {
		if v > 1 {
			ans = "NO"
		}
	}
	fmt.Println(ans)
}
