package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	for l := len(s)/2 - 1; ; l-- {
		if s[0:l] == s[l:2*l] {
			fmt.Println(2 * l)
			break
		}
	}
}
