package main

import "fmt"

func divs(n int) []int {
	ans := make([]int, 0)
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			ans = append(ans, i)
			n /= i
		}
	}
	if n != 1 {
		ans = append(ans, n)
	}
	return ans
}

func main() {
	var n int
	fmt.Scan(&n)

	l := make([]int, 0)
	for i := 2; i <= n; i++ {
		l = append(l, divs(i)...)
	}

	counter := make(map[int]int)
	for _, d := range l {
		counter[d]++
	}

	ans := 1
	for _, cnt := range counter {
		ans = (cnt + 1) * ans % 1000000007
	}
	fmt.Println(ans)
}
