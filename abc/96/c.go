package main

import "fmt"

func c(ss []string, y, x int) bool {
	if x < 0 || y < 0 || y >= len(ss) || x >= len(ss[y]) {
		return false
	}
	return ss[y][x] == '#'
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	ss := make([]string, h)
	for i := range ss {
		fmt.Scan(&ss[i])
	}

	for y := range ss {
		for x := range ss[y] {
			if ss[y][x] == '#' && !c(ss, y-1, x) && !c(ss, y+1, x) && !c(ss, y, x-1) && !c(ss, y, x+1) {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}
