// compare.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	now := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	end := time.Now()
	fmt.Println("cost time:", end.Sub(now))

	n_now := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	n_end := time.Now()
	fmt.Println("Join funcation cost time:", n_end.Sub(n_now))
}
