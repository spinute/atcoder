package main

import "fmt"

func check(s string) bool {
	var found [26]bool
	for _, r := range s {
		if found[int(r-'a')] {
			return false
		}
		found[int(r-'a')] = true
	}
	return true
}

func main() {
	var s string
	fmt.Scan(&s)
	if check(s) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
