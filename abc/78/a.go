package main

import "fmt"

func comp(a, b byte) string {
	if a < b {
		return "<"
	} else if a == b {
		return "="
	} else {
		return ">"
	}
}

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	fmt.Println(comp(a[0], b[0]))
}
