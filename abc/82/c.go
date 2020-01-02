package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	as := make([]int, n)
	for i := range as {
		fmt.Scan(&as[i])
	}
	counter := make(map[int]int)
	for _, a := range as {
		counter[a]++
	}

	ans := 0
	for a, cnt := range counter {
		if cnt >= a {
			ans += cnt - a
		} else {
			ans += cnt
		}
	}
	fmt.Println(ans)
}
