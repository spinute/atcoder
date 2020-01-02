package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	ss := make([]string, n)
	for i := range ss {
		fmt.Scan(&ss[i])
	}

	for i := 0; i < n; i++ {
		l := ""
		for j := n - 1; j >= 0; j-- {
			l += string(ss[j][i])
		}
		fmt.Println(l)
	}
}
