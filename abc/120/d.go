package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func wordScanner() *bufio.Scanner {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	return sc
}

func get_int(sc *bufio.Scanner) int64 {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return int64(i)
}

type entry struct {
	team int64
	size int64
}

func get_team(l []entry, i int64) int64 {
	ls := make([]int64, 0)
	for l[i].team != i {
		ls = append(ls, i)
		i = l[i].team
	}
	for _, v := range ls {
		l[v].team = i
	}
	return i
}

func main() {
	var n, m int64
	fmt.Scan(&n, &m)
	sc := wordScanner()
	as := make([]int64, m)
	bs := make([]int64, m)
	for i := range as {
		as[i] = get_int(sc) - 1
		bs[i] = get_int(sc) - 1
	}

	l := make([]entry, n)
	for i := range l {
		l[i].team = int64(i)
		l[i].size = 1
	}

	ans := make([]int64, m)
	ans[len(ans)-1] = n * (n - 1) / 2

	for i := len(ans) - 2; i >= 0; i-- {
		a := as[i+1]
		b := bs[i+1]
		ta := get_team(l, a)
		tb := get_team(l, b)
		if ta == tb {
			ans[i] = ans[i+1]
		} else {
			if ta < tb {
				l[tb].team = ta
				l[b].team = ta
			} else {
				l[ta].team = tb
				l[a].team = tb
			}
			sum := l[ta].size + l[tb].size
			diff := l[ta].size * l[tb].size
			ans[i] = ans[i+1] - diff
			l[a].size = sum
			l[b].size = sum
			l[ta].size = sum
			l[tb].size = sum
		}
	}

	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}
}
