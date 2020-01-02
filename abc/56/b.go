package main

import "fmt"

func main() {
	var w, a, b int
	fmt.Scan(&w, &a, &b)
	ans := 0
	if a+w < b {
		ans = b - a - w
	} else if b+w < a {
		ans = a - b - w
	} else {
		ans = 0
	}
	fmt.Println(ans)
}
