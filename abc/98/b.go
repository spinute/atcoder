package main

import "fmt"

func score(s1, s2 string) int {
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)
	for _, r := range s1 {
		m1[r]++
	}
	for _, r := range s2 {
		m2[r]++
	}

	cnt := 0
	for k := range m1 {
		if m2[k] > 0 {
			cnt++
		}
	}

	return cnt
}

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	max_score := score(s[:0], s[0:len(s)])
	for i := range s {
		if sc := score(s[:i], s[i:len(s)]); sc > max_score {
			max_score = sc
		}
	}
	fmt.Println(max_score)
}
