package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	as := make([]int, n)
	for i := range as {
		fmt.Scan(&as[i])
	}

	cnt := 0
	for _, a := range as {
		switch a % 6 {
		case 2:
			cnt += 1
		case 4:
			cnt += 1
		case 5:
			cnt += 2
		case 0:
			cnt += 3
		}
	}
	fmt.Println(cnt)
}
