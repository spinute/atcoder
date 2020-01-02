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
	max := 0
	for k, cnt := range counter {
		if cnt > 1 && k > max {
			max = k
		}
	}
	counter[max] -= 2
	second := 0
	for k, cnt := range counter {
		if cnt > 1 && k > second {
			second = k
		}
	}
	fmt.Println(max * second)
}
