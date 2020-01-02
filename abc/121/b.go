package main

import "fmt"

func main() {
	var n, m, c int
	fmt.Scan(&n, &m, &c)
	bs := make([]int, m)
	for i := range bs {
		fmt.Scan(&bs[i])
	}
	ass := make([][]int, n)
	for i := range ass {
		ass[i] = make([]int, m)
		for j := range ass[i] {
			fmt.Scan(&ass[i][j])
		}
	}

	ans := 0
	for _, as := range ass {
		sum := 0
		for j, a := range as {
			sum += a * bs[j]
		}
		if sum+c > 0 {
			ans++
		}
	}
	fmt.Println(ans)
}
