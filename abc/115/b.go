package main

import "fmt"

func main() {
	var n int
	var ins []int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var in int
		fmt.Scan(&in)
		ins = append(ins, in)
	}

	var max = ins[0]
	sum := 0
	for i := 0; i < n; i++ {
		sum += ins[i]
		if max < ins[i] {
			max = ins[i]
		}
	}

	fmt.Println(sum - max/2)
}
