package main

import "fmt"

func normalize(s string) []int {
	ret := make([]int, len(s))
	m := make(map[rune]int)
	cnt := 1
	for i, r := range s {
		if m[r] == 0 {
			m[r] = cnt
			cnt++
		}
		ret[i] = m[r]
	}
	return ret
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	if equal(normalize(s), normalize(t)) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
