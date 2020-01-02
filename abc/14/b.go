package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	as := make([]int, n)
	for i := range as {
		fmt.Scan(&as[i])
	}

	sum := 0
	for i := 0; i < n; i++ {
		if (1<<uint(i))&x != 0 {
			sum += as[i]
		}
	}
	fmt.Println(sum)
}
