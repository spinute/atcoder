package main

import "fmt"

func main() {
	var x, y int64
	fmt.Scan(&x, &y)
	l := 1
	for x*2 <= y {
		x *= 2
		l++
	}
	fmt.Println(l)
}
