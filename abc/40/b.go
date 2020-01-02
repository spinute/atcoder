package main

import "fmt"

func score(n, h, w int) int {
	return n - h*w + h - w
}

func main() {
	var n int
	fmt.Scan(&n)

	best := score(n, n, 1)
	for w := 1; w*w <= n; w++ {
		for h := w; h*w <= n; h++ {
			if v := score(n, h, w); v < best {
				best = v
			}
		}
	}
	fmt.Println(best)
}
