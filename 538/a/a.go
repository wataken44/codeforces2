
package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	var t string = "CODEFORCES"

	for m := 0; m < len(s); m++ {
		for n := m; n <= len(s); n++ {
			var s0 string = s[m:n]
			
		}
	}
	if len(t) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
