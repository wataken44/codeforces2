
package main

import (
	"fmt"
)

func main() {
	var n int
	var k int
	fmt.Scanf("%d %d", &n, &k)

	var a []int = make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	th := a[k-1]
	r := 0
	for i := 0; i < n; i++ {
		if a[i] > 0 && a[i] >= th {
			r++
		}
	}
	fmt.Println(r)
}
