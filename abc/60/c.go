package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

func get_int(sc *bufio.Scanner) int64 {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int64(i)
}

func main() {
	var n, t int64
	fmt.Scan(&n, &t)
	ts := make([]int64, n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := range ts {
		ts[i] = get_int(sc)
	}

	sum := int64(0)
	for i := 0; i < len(ts)-1; i++ {
		sum += min(ts[i+1]-ts[i], t)
	}
	sum += t
	fmt.Println(sum)
}
