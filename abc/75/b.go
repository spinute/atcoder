package main

import "fmt"

func in_range(ss []string, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(ss) && j < len(ss[i])
}

func count(ss []string, i, j int) int {
	cnt := 0
	for y := i - 1; y <= i+1; y++ {
		for x := j - 1; x <= j+1; x++ {
			if in_range(ss, y, x) && ss[y][x] == '#' {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	ss := make([]string, h)
	for i := range ss {
		fmt.Scan(&ss[i])
	}

	for i := range ss {
		for j := 0; j < len(ss[i]); j++ {
			if ss[i][j] == '#' {
				fmt.Print("#")
			} else {
				fmt.Print(count(ss, i, j))
			}
		}
		fmt.Println("")
	}
}
