
package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)
	s = strings.ToLower(s)
	
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a', 'o', 'y', 'e', 'u', 'i':
			// do nothing
		default:
			fmt.Printf(".%c", s[i])
		}
	}
	fmt.Printf("\n")
}
