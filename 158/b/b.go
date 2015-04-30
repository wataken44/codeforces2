
package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	var s []int = make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	
	sort.Sort(sort.IntSlice(s))

	var r int = 0
	var b int
	for len(s) > 0 {
		b, s = s[len(s)-1], s[:len(s)-1]
		var t int = b
		for len(s) > 0 {
			if s[0] + t <= 4 {
				t += s[0]
				s = s[1:]
			} else {
				break
			}
		}
		r++
	}
	fmt.Println(r)
}
