
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var o int = 0
	var t int = 0

	for i := 0; i < n; i++ {
		var p int
		fmt.Scan(&p)

		o += p
		t += 100
	}

	fmt.Printf("%.10f\n", float64(100*o) / float64(t))
}
