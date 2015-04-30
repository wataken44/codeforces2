
package main

import (
	"fmt"
	"strings"
)


func main() {
	var c []byte = make([]byte, 26, 26)
	for i := 0; i < 26; i++ {
		c[i] = 'a' + i
	}

	var n int
	fmt.Scan(&n)

	var s []string = make([]string, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	var l int = 1
	for {
		var v []string
		for i := 0; i < n; i++ {
			for k := 0; k < n - l + 1; k++ {
				v.Extend(s[i][k:k+l])
				fmt.Println(s[i][k:k+l])
			}
		}
		l += 1
		return
	}
}

