package main

import (
	"fmt"
	"strings"
)

type counter map[rune]int

func newCounter(s string) counter {
	ret := make(counter)
	for _, b := range s {
		ret[b]++
	}
	return ret
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func (c0 counter) intersect(c1 counter) counter {
	ret := make(counter)
	for k, v := range c0 {
		ret[k] = min(c1[k], v)
	}
	return ret
}

func main() {
	var n int
	fmt.Scan(&n)
	ss := make([]string, n)
	for i := range ss {
		fmt.Scan(&ss[i])
	}

	cntr := newCounter(ss[0])
	for _, s := range ss {
		cntr = cntr.intersect(newCounter(s))
	}

	ans := ""
	for c := 'a'; c <= 'z'; c++ {
		ans += strings.Repeat(string(c), cntr[c])
	}
	fmt.Println(ans)
}
