package dup

import (
	"bufio"
	"fmt"
	"io"
)

// FindDuplicates reads input lines and prints out which are duplicated with the
// nubmer of times.
func FindDuplicates(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)
	counts := make(map[string]int)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		counts[scanner.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Fprintf(output, "%d\t%s\n", n, line)
		}
	}
}
