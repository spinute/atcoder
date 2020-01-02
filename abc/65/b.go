package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	as := make([]int, n)
	for i := range as {
		fmt.Scan(&as[i])
	}

	pressed := make([]bool, n)
	for i, cnt := 0, 0; ; i, cnt = as[i]-1, cnt+1 {
		if i == 1 {
			fmt.Println(cnt)
			break
		}
		if pressed[i] {
			fmt.Println(-1)
			break
		}
		pressed[i] = true
	}
}
