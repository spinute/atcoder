package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	ans := ""
	for _, r := range s {
		switch r {
		case 'a', 'i', 'e', 'u', 'o':
		default:
			ans += string(r)
		}
	}
	fmt.Println(ans)
}
