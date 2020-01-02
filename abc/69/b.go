package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Printf("%c%d%c\n", s[0], len(s)-2, s[len(s)-1])
}
