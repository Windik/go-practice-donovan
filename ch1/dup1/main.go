package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	// Read lines from standard input
	// press Ctrl+D to finish input for MacOS
	// or Ctrl+Z to finish input for Windows
	for scanner.Scan() {
		counts[scanner.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
