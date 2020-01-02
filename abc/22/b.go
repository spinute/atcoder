package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	as := make([]int, n)
	for i := range as {
		fmt.Scan(&as[i])
	}

	m := make(map[int]int)
	for _, a := range as {
		m[a]++
	}
	cnt := 0
	for _, v := range m {
		cnt += v - 1
	}
	fmt.Println(cnt)
}
