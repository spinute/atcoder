package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func get_int(sc *bufio.Scanner) int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	as := make([]int, m)
	bs := make([]int, m)
	for i := range as {
		as[i] = get_int(sc)
		bs[i] = get_int(sc)
	}

	cand := make(map[int]bool, 0)
	for i, b := range bs {
		if b == n {
			cand[as[i]] = true
		}
	}
	for i, a := range as {
		if a == 1 && cand[bs[i]] == true {
			fmt.Println("POSSIBLE")
			return
		}
	}
	fmt.Println("IMPOSSIBLE")
}
