package main

import "fmt"

func main() {
	var n, t int
	fmt.Scan(&n, &t)
	as := make([]int, n)
	for i := range as {
		fmt.Scan(&as[i])
	}

	sum := 0
	prev := as[0]
	for _, a := range as {
		if v := a - prev; v > t {
			sum += t
		} else {
			sum += v
		}
		prev = a
	}
	fmt.Println(sum + t)
}
