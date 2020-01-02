package main

import "fmt"

func main() {
	var k, s int
	fmt.Scan(&k, &s)

	cnt := 0
	for x := 0; x <= k; x++ {
		for y := 0; y <= k; y++ {
			if z := s - x - y; 0 <= z && z <= k {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}
