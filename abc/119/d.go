package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func scanner() *bufio.Scanner {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	return sc
}

func get_int(sc *bufio.Scanner) int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	var a, b, q int
	fmt.Scan(&a, &b, &q)

	ss := make([]int, a)
	ts := make([]int, b)
	xs := make([]int, q)

	sc := scanner()
	for i := range ss {
		ss[i] = get_int(sc)
	}
	ss = append(ss, -100000000000)
	ss = append(ss, 100000000000)
	for i := range ts {
		ts[i] = get_int(sc)
	}
	ts = append(ts, -100000000000)
	ts = append(ts, 100000000000)
	for i := range xs {
		xs[i] = get_int(sc)
	}

	sort.Ints(ss)
	sort.Ints(ts)

	for _, x := range xs {
		si := sort.SearchInts(ss, x)
		ti := sort.SearchInts(ts, x)
		v0 := s1(ss[si], ts[ti], x)
		v1 := s1(ss[si-1], ts[ti-1], x)
		v2 := s2(ss[si-1], ts[ti], x)
		v3 := s2(ss[si], ts[ti-1], x)

		min := 100000000000
		for _, v := range []int{v0, v1, v2, v3} {
			if min > v {
				min = v
			}
		}
		fmt.Println(min)
	}
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func s1(a, b, x int) int {
	return max(diff(a, x), diff(b, x))
}

func s2(a, b, x int) int {
	d1 := diff(a, x)
	d2 := diff(x, b)
	return d1 + d2 + min(d1, d2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
