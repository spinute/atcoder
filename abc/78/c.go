package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	cost := (n-m)*100 + m*1900
	fmt.Println(cost * (1 << uint(m)))
}
