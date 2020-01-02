package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	ns := make([]int, n)
	for i := range ns {
		fmt.Scan(&ns[i])
	}

	max := 0
	ans := 0
	for _, n := range ns {
		if n >= max {
			max = n
			ans++
		}
	}
	fmt.Println(ans)
}
