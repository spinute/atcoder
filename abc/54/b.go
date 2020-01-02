package main

import "fmt"

func check(as, bs []string) bool {
	la := len(as)
	lb := len(bs)
	for j := 0; j <= la-lb; j++ {
	OUTER:
		for i := 0; i <= la-lb; i++ {
			for y := 0; y < lb; y++ {
				if as[j+y][i:i+lb] != bs[y] {
					continue OUTER
				}
			}
			return true
		}
	}
	return false
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	as := make([]string, n)
	bs := make([]string, m)
	for i := range as {
		fmt.Scan(&as[i])
	}
	for i := range bs {
		fmt.Scan(&bs[i])
	}

	ans := ""
	if check(as, bs) {
		ans = "Yes"
	} else {
		ans = "No"
	}
	fmt.Println(ans)
}
