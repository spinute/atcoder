package main

import (
	"fmt"
	"math/big"
)

func find_divisor(n int64) (bool, int64) {
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			return true, i
		}
	}
	return false, 0
}

func factors(n int64) []int64 {
	var ans []int64
	r := n
	for {
		found, i := find_divisor(r)
		if found {
			ans = append(ans, i)
			r /= i
		} else {
			if r != 1 {
				ans = append(ans, r)
			}
			break
		}
	}
	return ans
}

func count(vs []int64) map[int64]int64 {
	m := make(map[int64]int64)
	for _, v := range vs {
		m[v]++
	}
	return m
}

func fact(n int64) int64 {
	if n == 1 {
		return 1
	}
	return n * fact(n-1) % 1000000007
}

func comb(a, b int64) int64 {
	var z big.Int
	z.Binomial(a, b)
	z.Mod(&z, big.NewInt(1000000007))
	return int64(z.Uint64())
}

func main() {
	var n, m int64
	fmt.Scan(&n, &m)
	fs := factors(m)
	counter := count(fs)
	var ans int64 = 1
	for _, num := range counter {
		ans = ans * comb(num+n-1, num) % 1000000007
	}
	fmt.Println(ans)
}
