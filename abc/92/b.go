package main

import "fmt"

func main() {
	var n, d, x int
	fmt.Scan(&n, &d, &x)
	as := make([]int, n)
	for i := range as {
		fmt.Scan(&as[i])
	}

	var num_eaten int
	for _, a := range as {
		num_eaten += 1 + (d-1)/a
	}
	fmt.Println(x + num_eaten)
}
