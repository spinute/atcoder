package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	if s[0] == s[2] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
