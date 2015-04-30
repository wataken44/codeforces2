
package main

import (
	"fmt"
)

func parse(s string) (r map[byte]int) {
	r = make(map[byte]int)
	var c byte = 0
	for c = 'A'; c <= 'Z'; c++ {
		r[c] = 0
	}
	for c = 'a'; c <= 'z'; c++ {
		r[c] = 0
	}
	
	for i := 0; i < len(s); i++ {
		r[s[i]] += 1
	}
	return r
}

func main() {
	var s string
	var t string
	fmt.Scanln(&s)
	fmt.Scanln(&t)

	ms := parse(s)
	mt := parse(t)
    
	var y int = 0
	var w int = 0

	for k, _ := range ms {
		yy := ms[k]
		if yy > mt[k] {
			yy = mt[k]
		}
		y += yy
		ms[k] -= yy
		mt[k] -= yy
	}

	for k, _ := range ms {
		ww := ms[k]
		kk := k ^ 32
		if ww > mt[kk] {
			ww = mt[kk]
		}
		w += ww
		ms[k] -= ww
		mt[kk] -= ww
	}

	fmt.Printf("%d %d", y, w)
	fmt.Println("")
}
