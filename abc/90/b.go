package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	cnt := 0
	for i := 1; i < 1000; i++ {
		cand := i + i/10%10*1000 + i%10*10000
		if b >= cand && cand >= a {
			cnt++
		}
	}
	fmt.Println(cnt)
}
