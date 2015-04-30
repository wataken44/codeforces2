
package main

import (
	"fmt"
)

func is_quasi(i int) bool {
	for i > 0 {
		if i % 10 >= 2 {
			return false
		}
		i = i / 10
	}
	return true
}

func main() {
	var n int
	fmt.Scan(&n)

	var qb []int = make([]int, 0, 1 << 7)
	for i := 1; i <= n; i ++ {
		if is_quasi(i) {
			qb = append(qb, i);
		}
	}

	var dp []int = make([]int, n+1)
	var prev []int = make([]int, n+1)

	for i := 1; i < len(dp); i++ {
		dp[i] = n + 1
	}

	for a := 1; a < len(dp); a++ {
		for i := 0; i < len(qb); i++ {
			var q int = qb[i]
			if a < q {
				continue;
			}
			if dp[a] > dp[a - q] + 1 {
				dp[a] = dp[a - q] + 1
				prev[a] = q
			}
		}
	}
	fmt.Println(dp[n])
	for n > 0 {
		fmt.Print(prev[n])
		if n != prev[n] {
			fmt.Printf(" ")
		} else {
			fmt.Println("")
		}
		n -= prev[n]
	}
}
