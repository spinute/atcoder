package main

import "fmt"

func dump(m map[int]map[string]int, n int) {
	var cs = []string{"A", "C", "G", "T"}
	for i := 3; i <= n; i++ {
		for _, c1 := range cs {
			for _, c2 := range cs {
				for _, c3 := range cs {
					s := c1 + c2 + c3
					fmt.Println(s, m[i][s])
				}
			}
		}
	}
}

func dp(s string, n int, m map[int]map[string]int) {
	if s != "AGC" && s != "ACG" && s != "GAC" {
		s1 := "A" + s[0:2]
		s2 := "C" + s[0:2]
		s3 := "G" + s[0:2]
		s4 := "T" + s[0:2]
		m[n][s] = m[n-1][s1] + m[n-1][s2] + m[n-1][s3] + m[n-1][s4]
		if (s[0] == 'G' && s[2] == 'C') || (s[1:] == "GC") {
			m[n][s] -= m[n-1][s1]
		}
		m[n][s] %= 1000000007
	} else {
		m[n][s] = 0
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	var cs = []string{"A", "C", "G", "T"}

	memo := make(map[int]map[string]int)
	for i := 3; i <= n; i++ {
		memo[i] = make(map[string]int)
	}

	for _, c1 := range cs {
		for _, c2 := range cs {
			for _, c3 := range cs {
				s := c1 + c2 + c3
				if s != "AGC" && s != "ACG" && s != "GAC" {
					memo[3][s] = 1
				}
			}
		}
	}

	for i := 4; i <= n; i++ {
		for _, c1 := range cs {
			for _, c2 := range cs {
				for _, c3 := range cs {
					s := c1 + c2 + c3
					dp(s, i, memo)
				}
			}
		}
	}

	ans := 0
	for _, v := range memo[n] {
		ans += v
		ans %= 1000000007
	}
	fmt.Println(ans)
}
