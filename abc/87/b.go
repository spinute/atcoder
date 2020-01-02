package main

import "fmt"

func main() {
	var a, b, c, x int
	fmt.Scan(&a, &b, &c, &x)
	cnt := 0
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			p := x - 500*i - 100*j
			if p >= 0 && p%50 == 0 && p/50 <= c {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
