package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	as := make([]int, n)
	for i := range as {
		fmt.Scan(&as[i])
	}

	sum := 0
	for _, v := range as {
		sum += v
	}

	if sum%n != 0 {
		fmt.Println(-1)
	} else {
		average := sum / n
		cnt := 0
		for i, conn, num := 0, 0, 0; i < n; i++ {
			conn++
			num += as[i]
			if num == average*conn {
				conn = 0
				num = 0
			} else {
				cnt++
			}
		}
		fmt.Println(cnt)
	}
}
