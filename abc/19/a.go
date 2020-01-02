package main

import "fmt"

var p = fmt.Println

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	switch {
	case a <= b && b <= c,
		c <= b && b <= a:
		p(b)
	case b <= a && a <= c,
		c <= a && a <= b:
		p(a)
	default:
		p(c)
	}
}
