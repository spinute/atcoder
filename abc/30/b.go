package main

import "fmt"

func min(a, b float64) float64 {
	if a > b {
		return b
	}
	return a
}

func diff(a, b float64) float64 {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	v := diff(float64(6*m), float64(30*(n%12))+float64(m)/2)
	fmt.Println(min(v, 360-v))
}
