package main

import "fmt"

func remaining_xs(h, w int, ss []string) []int {
	xs := make([]int, 0)
OUTER:
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if ss[y][x] != '.' {
				xs = append(xs, x)
				continue OUTER
			}
		}
	}
	return xs
}

func remaining_ys(h, w int, ss []string) []int {
	ys := make([]int, 0)
OUTER:
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if ss[y][x] != '.' {
				ys = append(ys, y)
				continue OUTER
			}
		}
	}
	return ys
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	ss := make([]string, h)
	for i := range ss {
		fmt.Scan(&ss[i])
	}

	xs := remaining_xs(h, w, ss)
	ys := remaining_ys(h, w, ss)

	for _, y := range ys {
		for _, x := range xs {
			fmt.Printf("%c", ss[y][x])
		}
		fmt.Print("\n")
	}
}
