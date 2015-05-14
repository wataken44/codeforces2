
package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func mv(o int, d int) int {
	var r int = abs(d - o)
	r = min(r, abs(d - (o + 10)))
	r = min(r, abs(d + 10 - o))
	return r
}

func main() {
	var n int
	fmt.Scan(&n)
	var orig string
	fmt.Scan(&orig)
	var dest string
	fmt.Scan(&dest)

	var r int = 0
	for i := 0; i < n; i++ {
		var o int = int(orig[i] - '0')
		var d int = int(dest[i] - '0')
		r += mv(o, d)
	}

	fmt.Println(r)
}
