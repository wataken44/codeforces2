
package main

import (
	"fmt"
)

func main() {
	var a,b,c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	var rx []int = make([]int, 0, 1001)

	for x := 1; a * x < c; x++ {
		if (c - a * x) % b == 0 {
			rx = append(rx, x)
		}
	}
	fmt.Println(len(rx))
	for i := 0; i < len(rx); i++ {
		fmt.Printf("%d %d\n", rx[i], (c - a * rx[i]) / b)
	}
}
