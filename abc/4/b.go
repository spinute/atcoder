package main

import "fmt"

func main() {
	var css [4][4]string
	for i := range css {
		for j := range css[i] {
			fmt.Scan(&css[i][j])
		}
	}

	for i := 3; i >= 0; i-- {
		ss := css[i]
		fmt.Printf("%s %s %s %s\n", ss[3], ss[2], ss[1], ss[0])
	}
}
