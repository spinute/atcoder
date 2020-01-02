package main

import "fmt"

var p = fmt.Println
var s = fmt.Scan

func main() {
	var a, b, c int
	s(&a, &b, &c)
	if a == b {
		p(c)
	} else if b == c {
		p(a)
	} else {
		p(b)
	}
}
