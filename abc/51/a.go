package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	for i := range s {
		if i == 5 || i == 13 {
			fmt.Print(" ")
		} else {
			fmt.Printf("%c", s[i])
		}
	}
	fmt.Println("")
}
