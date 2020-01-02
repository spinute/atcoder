package main

import "fmt"

func main() {
	var n, x int

	fmt.Scan(&n, &x)

	var hs, ps []int
	hs = append(hs, 1)
	ps = append(ps, 1)
	for i := 1; i <= n; i++ {
		hs = append(hs, 2*hs[i-1]+3)
		ps = append(ps, 2*ps[i-1]+1)
	}

	sum := 0
	rem := x

	for level := n; rem > 0; level-- {
		if rem == hs[level] {
			sum += ps[level]
			break
		}

		rem -= 1
		if rem <= hs[level-1] {
			continue
		}

		sum += ps[level-1]
		rem -= hs[level-1]

		rem -= 1
		sum += 1

		if rem == hs[level]+1 {
			rem -= 1
		}
	}
	fmt.Println(sum)
}
