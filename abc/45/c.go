package main

import (
	"fmt"
	"strconv"
)

func toDigits(s string) []int {
	ret := make([]int, 0)
	for _, r := range s {
		i, _ := strconv.Atoi(string(r))
		ret = append(ret, i)
	}
	return ret
}

func calc(is []int, plus []bool) int {
	sum := 0
	n := is[0]
	for i, b := range plus {
		if b {
			sum += n
			n = is[i+1]
		} else {
			n = n*10 + is[i+1]
		}
	}
	return sum + n
}

func solve(is []int) int {
	l := len(is) - 1
	sum := 0
	for mask := 0; mask < 1<<uint(l); mask++ {
		ops := make([]bool, l)
		for i := 0; i < l; i++ {
			ops[i] = 1<<uint(i)&mask > 0
		}
		sum += calc(is, ops)
	}
	return sum
}

func main() {
	var s string
	fmt.Scan(&s)
	is := toDigits(s)
	fmt.Println(solve(is))
}
