package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	ret := solve(n)
	fmt.Println(ret)
}

func solve(n int) int {
	factors := make([]int, n+1)

	for i := 2; i <= n; i++ {
		curr := i
		for j := 2; j <= i; j++ {
			for curr%j == 0 {
				factors[j] += 1
				curr /= j
			}
		}
	}

	count := 0

	for i := 2; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			for k := j + 1; k <= n; k++ {
				if factors[i] >= 4 && factors[j] >= 4 && factors[k] >= 2 {
					count += 1
				}
				if factors[i] >= 4 && factors[j] >= 2 && factors[k] >= 4 {
					count += 1
				}
				if factors[i] >= 2 && factors[j] >= 4 && factors[k] >= 4 {
					count += 1
				}
			}
		}
	}
	for i := 2; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if factors[i] >= 24 && factors[j] >= 2 {
				count += 1
			}
			if factors[i] >= 2 && factors[j] >= 24 {
				count += 1
			}
			if factors[i] >= 14 && factors[j] >= 4 {
				count += 1
			}
			if factors[i] >= 4 && factors[j] >= 14 {
				count += 1
			}
		}
	}
	for i := 2; i <= n; i++ {
		if factors[i] >= 74 {
			count += 1
		}
	}

	return count
}
