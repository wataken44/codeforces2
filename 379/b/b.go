
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var a []int = make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var f int = 0
	var p int = 0
	var b bool = true

	for f < n {
		var c byte
		switch {
		case p < f:
			c = 'R'
			p += 1
			b = true
		case p == f:
			if a[p] == 0 {
				f += 1
			}
			if b && a[p] > 0 {
				c = 'P'
				a[p] -= 1
				b = false
				if a[p] == 0 {
					f += 1
				}
			} else {
				if p != n - 1 {
					c = 'R'
					p += 1
				} else {
					c = 'L'
					p -= 1
				}
				b = true
			}
		case p > f:
			if b && a[p] > 0 {
				c = 'P'
				a[p] -= 1
				b = false
			} else {
				if p != 0 {
					c = 'L'
					p -= 1
				} else {
					// never reach
					c = 'R'
					p += 1
				}
				b = true
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println("")
}
