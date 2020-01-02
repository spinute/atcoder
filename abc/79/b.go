package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	ls := make([]int, n+1)
	ls[0] = 2
	ls[1] = 1
	for i := 2; i < len(ls); i++ {
		ls[i] = ls[i-1] + ls[i-2]
	}
	fmt.Println(ls[n])
}
