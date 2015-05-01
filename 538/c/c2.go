
package main

import (
	"fmt"
)

func abs(x int) int {
	return max(x, -x)
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func minmax(x int, y int) (int, int) {
	if x < y {
		return x, y
	}
	return y, x
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	fmt.Printf("%d %d\n", n, m)
	
	var d []int = make([]int, m + 1)
	var h []int = make([]int, m + 1)

	d[0] = 0
	h[0] = 0
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d", &d[i + 1], &h[i + 1])
		fmt.Printf("%d %d\n", d[i + 1], h[i + 1])
	}

	var ret int = 0
	for i := 1; i < len(d); i++ {
		var dd, dh int
		dd = d[i] - d[i - 1]
		dh = h[i] - h[i - 1]
		if abs(dd) < abs(dh) {
			fmt.Println("IMPOSSIBLE")
			return
		} else {
			ret = max(ret, h[i])
			ret = max(ret, min(h[i], h[i-1]) + (abs(dd) + abs(dh)) / 2)
		}
	}
	fmt.Println(ret)
}
