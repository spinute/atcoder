package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	as := make([]int, m)
	bs := make([]int, m)
	for i := range as {
		fmt.Scan(&as[i])
		fmt.Scan(&bs[i])
	}
	counter := make(map[int]int)
	for _, v := range as {
		counter[v]++
	}
	for _, v := range bs {
		counter[v]++
	}
	for i := 1; i <= n; i++ {
		fmt.Println(counter[i])
	}
}
