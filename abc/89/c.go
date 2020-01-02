package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	ss := make([]string, n)
	for i := range ss {
		fmt.Scan(&ss[i])
	}

	m := make(map[byte]int)
	for _, s := range ss {
		m[s[0]]++
	}

	cs := []byte{'M', 'A', 'R', 'C', 'H'}

	ans := int64(0)
	for i := range cs {
		for j := i + 1; j < len(cs); j++ {
			for k := j + 1; k < len(cs); k++ {
				ans += int64(m[cs[i]]) * int64(m[cs[j]]) * int64(m[cs[k]])
			}
		}
	}
	fmt.Println(ans)
}
