
package main

import (
	"fmt"
)

func lucky(x int, n int) int {
	var r int = 0
	for i := (n - 1); i >= 0; i-- {
		r *= 10
		if ((x >> uint(i)) & 1) == 1 {
			r += 7
		} else {
			r += 4
		}
	}
	return r
}

func main() {
	var n int
	fmt.Scan(&n)

	var c int = 1
	for i := 1; i <= 9; i++ {
		for k := 0; k < (1 << uint(i)); k++ {
			if n == lucky(k, i) {
				fmt.Println(c)
				break
			}
			c += 1
		}
	}
}
