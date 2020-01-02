package main

import "fmt"

func main() {
	var a, b, c int64
	fmt.Scan(&a, &b, &c)

	fmt.Println(a * b % 1000000007 * c % 1000000007)
}
