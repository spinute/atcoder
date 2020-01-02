package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fss := make([][10]int, n)
	for i := range fss {
		for j := range fss[i] {
			fmt.Scan(&fss[i][j])
		}
	}
	pss := make([][11]int, n)
	for i := range pss {
		for j := range pss[i] {
			fmt.Scan(&pss[i][j])
		}
	}

	max := -1000000009
	for mask := 1; mask < 1<<10; mask++ {
		sum := 0
		for i, fs := range fss {
			cnt := 0
			for j := 0; j < 10; j++ {
				if mask&(1<<uint(j)) > 0 && fs[j] == 1 {
					cnt++
				}
			}
			sum += pss[i][cnt]
		}
		if sum > max {
			max = sum
		}
	}
	fmt.Println(max)
}
