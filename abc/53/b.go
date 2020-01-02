package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	first_a_ind := -1
	last_z_ind := -1
	for i, r := range s {
		if r == 'A' && first_a_ind == -1 {
			first_a_ind = i
		}
		if r == 'Z' {
			last_z_ind = i
		}
	}

	fmt.Println(last_z_ind - first_a_ind + 1)
}
