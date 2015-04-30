
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	switch {
	case n <= 2:
		fmt.Println("1")
		fmt.Println("1")
	case n == 3:
		fmt.Println("2")
		fmt.Println("1 3")
	default:
		fmt.Println(n)
		a := 2
		for i := 0; i < n; i++ {
			fmt.Print(a)
			if i != n - 1 {
				fmt.Print(" ")
			} else {
				fmt.Println("")
			}
			a += 2
			if a > n {
				a = 1
			}
		}
	}
	
}
