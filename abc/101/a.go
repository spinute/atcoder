package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var i int
	for _, c := range s {
		if c == '+' {
			i++
		} else {
			i--
		}
	}
	fmt.Println(i)
}
