package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	ss := make([]int, n)
	for i := range ss {
		fmt.Scan(&ss[i])
	}

	sum := 0
	for _, s := range ss {
		sum += s
	}

	min := 101
	for _, s := range ss {
		if s < min && s%10 != 0 {
			min = s
		}
	}

	if sum%10 == 0 {
		if min == 101 {
			fmt.Println(0)
		} else {
			fmt.Println(sum - min)
		}
	} else {
		fmt.Println(sum)
	}
}
