
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var q string
	fmt.Scan(&q)

	var f [26]int

	for i := 0; i < len(q); i++ {
		var ci int = int(q[i] - 'a')
		f[ci] = 1
	}

	var c int = 0
	for i := 0; i < len(f); i++ {
		c += f[i]
	}

	if c < n {
		fmt.Println("NO")
		return
	}

	fmt.Print("YES")
	var g [26]int
	for i := 0; i < len(q); i++ {
		var ci int = int(q[i] - 'a')
		if (n > 0) && (g[ci] == 0) {
			fmt.Printf("\n%c", q[i])
			g[ci] = 1
			n--			
		} else {
			fmt.Printf("%c", q[i])
		}
	}
	fmt.Println("")
}
