package main

import "fmt"

func main() {
	var s string
	var i int
	fmt.Scan(&s, &i)
	fmt.Printf("%c\n", s[i-1])
}
