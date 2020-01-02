package main

import "fmt"

var p = fmt.Println

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	if len(s) > len(t) {
		p(s)
	} else {
		p(t)
	}
}
