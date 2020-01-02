package main

import "fmt"

func main() {
	var a, b, c, k, s, t int
	fmt.Scan(&a, &b, &c, &k, &s, &t)
	sum := a*s + b*t
	if s+t >= k {
		sum -= (s + t) * c
	}
	fmt.Println(sum)
}
