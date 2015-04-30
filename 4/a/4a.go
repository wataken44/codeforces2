package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		w, _ := strconv.Atoi(scanner.Text())
		if (w > 2) && (w%2 == 0) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
