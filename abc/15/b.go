package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	as := make([]int, n)
	for i := range as {
		fmt.Scan(&as[i])
	}

	cnt := 0
	sum := 0
	for _, v := range as {
		if v > 0 {
			cnt++
			sum += v
		}
	}

	fmt.Println((sum-1)/cnt + 1)
}
