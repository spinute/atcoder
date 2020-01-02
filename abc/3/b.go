package main

import "fmt"

func ok(r0, r byte) bool {
	if r0 == '@' {
		switch r {
		case 'a', 't', 'c', 'o', 'd', 'e', 'r', '@':
			return true
		}
		return false
	}
	if r == '@' {
		switch r0 {
		case 'a', 't', 'c', 'o', 'd', 'e', 'r', '@':
			return true
		}
		return false
	}
	return false
}

func check(s, t string) bool {
	for i := range s {
		if s[i] != t[i] && !ok(s[i], t[i]) {
			return false
		}
	}
	return true
}

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	if check(s, t) {
		fmt.Println("You can win")
	} else {
		fmt.Println("You will lose")
	}
}
