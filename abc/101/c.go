package main

import "fmt"

func index(val int, vs []int) int {
	for i, v := range vs {
		if v == val {
			return i
		}
	}
	return -1
}

func div_roundup(a, b int) int {
	if a == 0 {
		return 0
	}
	return (a-1)/b + 1
}

func main() {
	var N, K int
	fmt.Scan(&N, &K)
	as := make([]int, N)
	for i := range as {
		fmt.Scan(&as[i])
	}

	i := index(1, as)
	l := i
	r := len(as) - 1 - i
	fmt.Println(l/(K-1) + r/(K-1) + div_roundup(l%(K-1)+r%(K-1), K-1))
}
