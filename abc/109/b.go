package main

import "fmt"

func is_unique(ss []string) bool {
	smap := make(map[string]int)
	for _, v := range ss {
		smap[v]++
	}
	for _, v := range smap {
		if v > 1 {
			return false
		}
	}
	return true
}

func is_connectable(ss []string) bool {
	for i := 0; i < len(ss)-1; i++ {
		from := ss[i]
		to := ss[i+1]
		if from[len(from)-1] != to[0] {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Scan(&n)
	ss := make([]string, n)
	for i := range ss {
		fmt.Scan(&ss[i])
	}
	if is_unique(ss) && is_connectable(ss) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
