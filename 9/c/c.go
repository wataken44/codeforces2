
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	var m int = 1
	var r int = 0
	for {
		var s string = strconv.FormatInt(int64(m), 2)
		k, _ := strconv.Atoi(s)
		if k <= n {
			r++
		}
		if k > n {
			break
		}
		m++
	}
	fmt.Println(r)
}
