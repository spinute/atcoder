package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	xs := make([]float64, n)
	us := make([]string, n)
	for i := range xs {
		fmt.Scan(&xs[i], &us[i])
	}

	ans := 0.0
	for i, x := range xs {
		if us[i] == "BTC" {
			x *= 380000.0
		}
		ans += x
	}

	fmt.Printf("%f\n", ans)
}
