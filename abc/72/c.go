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
		counter[a-1]++
		counter[a]++
		counter[a+1]++
	}
	max := 0
	for _, v := range counter {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
}
