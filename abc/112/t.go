package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(s[0] - 'A' + 1)
}
