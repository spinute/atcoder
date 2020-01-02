package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	if 2*n >= m {
		fmt.Println(m / 2)
	} else {
		fmt.Println(n + (m-2*n)/4)
	}
}
