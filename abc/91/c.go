package main

import (
	"fmt"
	"sort"
)

type blue struct {
	x, y int
}
type blues []blue

func (vs blues) Len() int {
	return len(vs)
}
func (vs blues) Swap(i, j int) {
	vs[i], vs[j] = vs[j], vs[i]
}
func (vs blues) Less(i, j int) bool {
	return vs[i].x < vs[j].x
}

type red struct {
	x, y int
}
type reds []red

func (vs reds) Len() int {
	return len(vs)
}
func (vs reds) Swap(i, j int) {
	vs[i], vs[j] = vs[j], vs[i]
}
func (vs reds) Less(i, j int) bool {
	return vs[i].y < vs[j].y
}

func main() {
	var n int
	fmt.Scan(&n)
	rs := make(reds, n)
	bs := make(blues, n)
	for i := range rs {
		fmt.Scan(&rs[i].x, &rs[i].y)
	}
	for i := range bs {
		fmt.Scan(&bs[i].x, &bs[i].y)
	}

	sort.Sort(sort.Reverse(rs))
	sort.Sort(bs)

	cnt := 0
	used := make([]bool, n)
	for _, b := range bs {
		for i, r := range rs {
			if b.x > r.x && b.y > r.y && !used[i] {
				cnt++
				used[i] = true
				break
			}
		}
	}

	fmt.Println(cnt)
}
