
package main

import (
	"fmt"
	"strconv"
)

func isLucky(n int) bool {
	for n > 0 {
		m := n % 10
		if !(m == 4 || m == 7) {
			return false
		}
		n /= 10
	}
	return true
}

func main() {
	var s string
	var n int
	fmt.Scanf("%s", &s)
	n, _ = strconv.Atoi(s)

	var ok bool = false
	for i := 4; i <= n; i++ {
		if n % i == 0 && isLucky(i) {
			ok = true
		}
	}
	if ok {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
