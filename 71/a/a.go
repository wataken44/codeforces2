
package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scanf("%s", &s)
		s = strings.TrimRight(s, " \r\n")
		if len(s) <= 10 {
			fmt.Println(s)
		} else {
			fmt.Printf("%c%d%c\n", s[0], len(s)-2, s[len(s)-1])
		}
	}
}
