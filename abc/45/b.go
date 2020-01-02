package main

import "fmt"

func pi2p(i int) string {
	return string(int('A') + i)
}

func p2pi(p byte) int {
	return int(p - 'a')
}

func main() {
	var ss [3]string
	fmt.Scan(&ss[0], &ss[1], &ss[2])
	pi := 0
	is := make([]int, 3)
	for {
		s := ss[pi]
		i := is[pi]
		if i < len(s) {
			is[pi]++
			pi = p2pi(s[i])
		} else {
			fmt.Println(pi2p(pi))
			break
		}
	}
}
