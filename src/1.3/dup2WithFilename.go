// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// map[filename][line]count
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines("<stdin>", os.Stdin, counts)
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(filename, f, counts)
			f.Close()
		}
	}
	for filename, lineMap := range counts {
		for line, n := range lineMap {
			if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", filename, n, line)
			}
		}
	}
}

func countLines(filename string, f *os.File, counts map[string]map[string]int) {
	// 安全初始化 map 的标准写法，用于处理嵌套 map 的情况
	if counts[filename] == nil {
		counts[filename] = make(map[string]int)
	}

	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[filename][line]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
