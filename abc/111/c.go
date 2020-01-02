package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func find_candidates(m map[int]int) (int, int, int, int) {
	f, fi, s, si := 0, 0, 0, 0
	for i, cnt := range m {
		if cnt > f {
			s = f
			si = fi
			f = cnt
			fi = i
		} else if cnt > s {
			s = cnt
			si = i
		}
	}
	return f, fi, s, si
}

func main() {
	var n int
	fmt.Scan(&n)
	cntrA := make(map[int]int)
	cntrB := make(map[int]int)
	for i := 0; i < n; i += 2 {
		var a, b int
		fmt.Scan(&a, &b)
		cntrA[a]++
		cntrB[b]++
	}

	fa, fai, sa, sai := find_candidates(cntrA)
	fb, fbi, sb, sbi := find_candidates(cntrB)

	_ = sai
	_ = sbi
	if fai != fbi {
		fmt.Println(n - fa - fb)
	} else {
		fmt.Println(min(n-sa-fb, n-fa-sb))
	}
}
