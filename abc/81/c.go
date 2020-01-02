package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	as := make([]int, n)
	for i := range as {
		fmt.Scan(&as[i])
	}
	counter := make(map[int]int)
	for _, a := range as {
		counter[a]++
	}
	if l := len(counter); k >= l {
		fmt.Println(0)
	} else {
		num_changes := l - k
		nums := make([]int, 0)
		for _, v := range counter {
			nums = append(nums, v)
		}
		sort.Ints(nums)
		sum := 0
		for i := 0; i < num_changes; i++ {
			sum += nums[i]
		}
		fmt.Println(sum)
	}
}
