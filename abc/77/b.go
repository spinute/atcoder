package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	i := 1
	for i*i <= n {
		i++
	}
	fmt.Println((i - 1) * (i - 1))
}
