package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	n := b - a - 1
	fmt.Println(n*(n+1)/2 - a)
}
