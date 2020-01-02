package main

import "fmt"

func last(s string) byte {
	return s[len(s)-1]
}

func main() {
	var a, b, c string
	fmt.Scan(&a, &b, &c)
	if last(a) == b[0] && last(b) == c[0] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
