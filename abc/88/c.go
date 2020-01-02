package main

import "fmt"

func main() {
	var css [3][3]int
	for i, cs := range css {
		for j := range cs {
			fmt.Scan(&css[i][j])
		}
	}

	as := make([]int, 3)
	bs := make([]int, 3)
	as[0] = 0
	as[1] = css[1][0] - css[0][0]
	as[2] = css[2][0] - css[0][0]
	bs[0] = css[0][0]
	bs[1] = css[0][1]
	bs[2] = css[0][2]

	for i, cs := range css {
		for j, c := range cs {
			if c != as[i]+bs[j] {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}
