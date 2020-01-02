package main

import "fmt"

func main() {
	var l, h, n int
	fmt.Scan(&l, &h, &n)
	as := make([]int, n)
	for i := range as {
		fmt.Scan(&as[i])
	}

	for _, v := range as {
		if v > h {
			fmt.Println(-1)
		} else if v >= l {
			fmt.Println(0)
		} else {
			fmt.Println(l - v)
		}
	}
}
