package main

import (
	"fmt"
	"strings"
)

func main() {
	var sx, sy, tx, ty int
	fmt.Scan(&sx, &sy, &tx, &ty)
	dx := tx - sx
	dy := ty - sy
	ans := ""
	ans += strings.Repeat("U", dy)
	ans += strings.Repeat("R", dx)
	ans += strings.Repeat("D", dy)
	ans += strings.Repeat("L", dx)
	ans += "L"
	ans += strings.Repeat("U", dy+1)
	ans += strings.Repeat("R", dx+1)
	ans += "D"
	ans += "R"
	ans += strings.Repeat("D", dy+1)
	ans += strings.Repeat("L", dx+1)
	ans += "U"
	fmt.Println(ans)
}
