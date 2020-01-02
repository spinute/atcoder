package main

import "fmt"

func score(as []uint64, mask uint64) uint64 {
	sum := uint64(0)
	for _, a := range as {
		sum += (a ^ mask) & mask
	}
	return sum
}

func max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func get_mask(n uint64) uint64 {
	ans := uint64(1) << 60
	for ans&n == 0 {
		ans >>= 1
	}
	return ans
}

func rev(as []uint64) []uint64 {
	bs := make([]uint64, len(as))
	for i := range bs {
		bs[i] = ^as[i]
	}
	return bs
}

func best(as []uint64, bits uint64) uint64 {
	sum := uint64(0)
	bs := rev(as)
	var i uint64
	for i = 1; bits&i != 0; i <<= 1 {
		sum += max(score(as, i), score(bs, i))
	}
	return sum
}

func solve(k uint64, as []uint64) uint64 {
	if k == 0 {
		return 0
	}
	mask := get_mask(k)
	v0 := score(rev(as), mask) + best(as, mask-1)
	s := score(as, mask)
	for m := mask >> 1; ; m >>= 1 {
		if m == 0 || m&k != 0 {
			break
		}
		s += score(rev(as), m)
	}
	v1 := s + solve(^mask&k, as)
	return max(v0, v1)
}

func main() {
	var n, k uint64
	fmt.Scan(&n, &k)
	as := make([]uint64, n)
	for i := range as {
		fmt.Scan(&as[i])
	}
	if k == 0 {
		s := uint64(0)
		for _, a := range as {
			s += 0 ^ a
		}
		fmt.Println(s)
	} else {
		fmt.Println(solve(k, as))
	}
}
